package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 uint64; F2 int }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS9 struct { F0 int64; F1 int; F2 uint16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int; F1 uint64 }
type GenS12 struct { F0 int64; F1 bool }
type GenS13 struct { F0 int8; F1 uint32; F2 int64 }
type GenS14 struct { F0 uint64; F1 int }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint; F1 int8; F2 uint32 }
type GenS18 struct { F0 uint32; F1 uint32 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint64; F1 bool }
type GenS21 struct { F0 int8; F1 int64; F2 uint32 }
type GenS22 struct { F0 bool; F1 uint8 }
type GenS23 struct { F0 uint8 }
type GenS24 struct { F0 int; F1 int8 }
type GenS25 struct { F0 string; F1 int; F2 uint32 }
func genFn0(p0 int64, p1 int) bool {
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: 42, F1: -1}
        v8 := []int{fi6, (p1 / ((fi6) | 1)), fi6}
        switch "" {
        case "foo bar":
            var v9 int8 = 1
            _ = v9
        case "":
            switch "\n\t" {
            case "":
                var v10 uint = 469
                _ = v10
            case "foo bar":
                for fi11 := 0; fi11 < 17; fi11++ {
                    switch fi6 {
                    case 128:
                        var v12 uint = 255
                        _ = v12
                    case -175:
                        if false {
                            if len(v8) > 0 {
                                v8[0] = 32767
                            }
                            if len(v8) > 0 {
                                v8[0] = (p1 ^ (p1 << ((18446744073709551615) & 63)))
                            }
                        }
                    case 1:
                        if false {
                            var v13 int32 = 212
                            if len(v8) > 0 {
                                v8[0] = p1
                            }
                            v14 := GenS2{F0: int8(p1)}
                            _ = v13
                            _ = v14
                        }
                        if false {
                            switch 876 {
                            case 128:
                                var v15 int32 = 42
                                _ = v15
                            case 0:
                                var v16 float32 = (0.1 * 1.5)
                                _ = v16
                            }
                            p1 += (p1 * 1)
                            var v17 uint32 = 42
                            _ = v17
                        }
                        if (2147483647 > 1) {
                            if true {
                                var v18 bool = false
                                if (!(!v18)) {
                                    if (fi6 == (p1 + p1)) {
                                        p1 = fi11
                                        if v18 {
                                            if (-1 == 127) {
                                                v19 := GenS3{F0: 2}
                                                var v20 uint16 = 2
                                                p1 = (p1 * (fi11 | (p1 & 4294967295)))
                                                _ = v19
                                                _ = v20
                                            } else {
                                                if false {
                                                    if false {
                                                        fmt.Println(fi6)
                                                    } else {
                                                        var v21 int32 = 1
                                                        _ = v21
                                                    }
                                                    v8 = append(v8, (-(fi6)))
                                                    if len(v8) > 0 {
                                                        v8[0] = fi11
                                                    }
                                                }
                                            }
                                            v22 := GenS4{F0: (fi11 << ((255) & 63)), F1: 1}
                                            _ = v22
                                        }
                                        switch "" {
                                        case "":
                                            switch 128 {
                                            case 127:
                                                switch p1 {
                                                case 0:
                                                    v23 := GenS4{F0: fi6, F1: 0}
                                                    _ = v23
                                                default:
                                                    v18 = (1 != 128)
                                                }
                                            default:
                                                switch "" {
                                                case "":
                                                    v18 = v18
                                                default:
                                                    if false {
                                                        p0 += p0
                                                        p1 -= fi6
                                                    } else {
                                                        if len(v8) > 0 {
                                                            v8[0] = (^(fi11))
                                                        }
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            fmt.Println(fi6)
                                        }
                                    }
                                    v24 := GenS2{F0: -85}
                                    fmt.Println(v18, p1)
                                    v25 := GenS5{F0: -1}
                                    v18 = v18
                                    var v26 uint64 = 0
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                                _ = v18
                            } else {
                                var v27 int16 = 40
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                _ = v27
                            }
                        }
                    default:
                        var v28 int16 = 42
                        v29 := GenS2{F0: -52}
                        _ = v28
                        _ = v29
                    }
                    switch "a" {
                    case "a":
                        var v30 int64 = (^(p0))
                        _ = v30
                    case "\n\t":
                        if ((!true) || true) {
                            p1 = fi11
                            p1 = (p1 | p1)
                            if len(v8) > 0 {
                                v8[0] = fi11
                            }
                        }
                    case "foo bar":
                        v31 := GenS6{F0: 0, F1: 42, F2: fi6}
                        _ = v31
                    default:
                        if true {
                            var v32 int = -453
                            var v33 uint16 = 2
                            v34 := GenS1{F0: v33, F1: 127}
                            _ = v32
                            _ = v33
                            _ = v34
                        } else {
                            if (float64((0.1 * (-2.25 - ((0.1 + 0.0) / 1.5)))) >= -2.25) {
                                var v35 float64 = float64(0.1)
                                var v36 float64 = 3.14159
                                switch 4294967295 {
                                case 32767:
                                    if len(v8) > 0 {
                                        v8[0] = 32767
                                    }
                                    p1 -= -1
                                case 127:
                                    var v37 float32 = -2.25
                                    _ = v37
                                case -244:
                                    p1 += 32767
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        switch "a" {
                                        case "a":
                                            _ = len(v8)
                                        }
                                    case "foo bar":
                                        v8 = append(v8, 9223372036854775807)
                                    case "hello":
                                        v38 := GenS7{F0: 42, F1: (p1 - -641)}
                                        _ = v38
                                    default:
                                        if (0 < 0) {
                                            if ((p1 % ((32768) | 1)) != fi6) {
                                                v8 = append(v8, fi6)
                                                var v39 int8 = -42
                                                _ = v39
                                            } else {
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(fi6, v36)
                                                case "":
                                                    if (42 < 2) {
                                                        var v40 uint16 = 647
                                                        v8 = append(v8, (p1 | 475))
                                                        v35 -= 0.0
                                                        if (65535 < 65535) {
                                                            v41 := GenS4{F0: fi11, F1: 0}
                                                            if false {
                                                                var v42 int16 = 0
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                                _ = v42
                                                            } else {
                                                                switch 4294967295 {
                                                                case -250:
                                                                    if len(v8) > 0 {
                                                                        v8[0] = fi6
                                                                    }
                                                                case 996:
                                                                    if len(v8) > 0 {
                                                                        v8[0] = fi11
                                                                    }
                                                                default:
                                                                    var v43 int32 = -1
                                                                    _ = v43
                                                                }
                                                            }
                                                            v44 := GenS8{F0: p0, F1: 255, F2: 65535}
                                                            _ = v41
                                                            _ = v44
                                                        }
                                                        _ = v40
                                                    } else {
                                                        if false {
                                                            if len(v8) > 0 {
                                                                v8[0] = (p1 ^ fi6)
                                                            }
                                                            switch "a" {
                                                            case "hello":
                                                                v35 += v36
                                                            case "a":
                                                                if true {
                                                                    v45 := GenS7{F0: 255, F1: (fi11 << ((0) & 63))}
                                                                    var v46 int8 = 1
                                                                    _ = v45
                                                                    _ = v46
                                                                }
                                                            case "\n\t":
                                                                if true {
                                                                    if len(v8) > 0 {
                                                                        v8[0] = (-(p1))
                                                                    }
                                                                    p1 = 4294967295
                                                                    v47 := GenS9{F0: (p0 ^ (p0 | p0)), F1: fi6, F2: 65535}
                                                                    _ = v47
                                                                } else {
                                                                    v35 = 0.1
                                                                    v48 := GenS10{F0: p0}
                                                                    var v49 uint32 = 42
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                            default:
                                                                switch "hello" {
                                                                case "a":
                                                                    v8 = append(v8, fi6)
                                                                case "\n\t":
                                                                    if false {
                                                                        switch (-(fi11)) {
                                                                        case -9223372036854775808:
                                                                            v8 = append(v8, p1)
                                                                        case 2147483647:
                                                                            var v50 int64 = -1
                                                                            _ = v50
                                                                        case 32767:
                                                                            var v51 float32 = ((1e10 * 0.1) + (((0.1 + -2.25) / (1e10 + 1e10)) / 1.0))
                                                                            _ = v51
                                                                        }
                                                                        v52 := GenS3{F0: 80}
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                        _ = v52
                                                                    } else {
                                                                        fmt.Println(p1, fi6, fi11)
                                                                        p1 -= fi11
                                                                        var v53 int8 = 92
                                                                        v8 = append(v8, -1)
                                                                        p1 = 664
                                                                        var v54 uint16 = 42
                                                                        _ = v53
                                                                        _ = v54
                                                                    }
                                                                }
                                                            }
                                                            v55 := GenS11{F0: fi6, F1: 4294967295}
                                                            _ = v55
                                                        } else {
                                                            v56 := GenS2{F0: -128}
                                                            var v57 uint = 96
                                                            if true {
                                                                v36 = (3.14159 - (v36 - (1.5 + v36)))
                                                            } else {
                                                                v58 := GenS8{F0: p0, F1: 65535, F2: 429}
                                                                v59 := GenS3{F0: 132}
                                                                if ((float64(1e10) / v36) > v36) {
                                                                    v60 := GenS12{F0: (p0 & (p0 + -720)), F1: ("" >= "hello")}
                                                                    _ = v60
                                                                } else {
                                                                    var v61 uint8 = 127
                                                                    _ = len(v8)
                                                                    _ = v61
                                                                }
                                                                fmt.Println(v35)
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        var v62 int8 = -96
                                                        _ = v62
                                                    }
                                                case "\n\t":
                                                    var v63 uint64 = 1
                                                    _ = v63
                                                }
                                            }
                                            var v64 float64 = 3.14159
                                            if (int32(fi6) != 490) {
                                                if true {
                                                    v65 := GenS4{F0: int(p0), F1: 42}
                                                    var v66 uint64 = 678
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    if true {
                                                        if false {
                                                            v35 = (1.5 - (((v64 / float64(((3.14159 + (float32((float64((3.14159 / (1.5 + 3.14159))) + (float64(3.14159) / v35))) / (float32((0.0 * ((3.14159 / float64(0.1)) - v64))) / ((1e10 - ((0.0 / (1.5 + -2.25)) + -2.25)) * 1e10)))) * ((3.14159 / 0.1) * 1e10)))) * v35) - 0.0))
                                                        } else {
                                                            _ = len(v8)
                                                        }
                                                        if len(v8) > 0 {
                                                            v8[0] = (fi6 % ((42) | 1))
                                                        }
                                                    }
                                                    var v67 int32 = 0
                                                    _ = v67
                                                }
                                                v68 := GenS13{F0: 1, F1: 2147483648, F2: (p0 << ((4294967295) & 63))}
                                                _ = v68
                                            }
                                            _ = v64
                                        } else {
                                            fmt.Println(p1, v36)
                                            var v69 int64 = 42
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            _ = v69
                                        }
                                        p1 += int(p0)
                                        var v70 float64 = float64(((1e10 / 0.1) * float32(((v35 + ((float64((1.5 * (0.0 * 3.14159))) - v35) * v36)) - ((float64(1e10) / v36) / -2.25)))))
                                        _ = v70
                                    }
                                }
                                if ((fi11 > (p1 * fi11)) || (false && (127 == 43))) {
                                    var v71 int32 = 0
                                    if len(v8) > 0 {
                                        v8[0] = (fi11 ^ p1)
                                    }
                                    _ = v71
                                } else {
                                    p0 = (p0 % ((-9223372036854775808) | 1))
                                    var v72 bool = (true || false)
                                    _ = v72
                                }
                                switch (fi11 - 1) {
                                case 887:
                                    var v73 uint32 = 2147483648
                                    _ = v73
                                case -2147483648:
                                    p1 = 2147483647
                                case 1:
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                default:
                                    if true {
                                        var v74 uint8 = 42
                                        v36 = v36
                                        var v75 int16 = 1000
                                        _ = v74
                                        _ = v75
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            v8 = append(v8, (p1 * 42))
                                        case "a":
                                            v76 := GenS5{F0: 0}
                                            if true {
                                                if true {
                                                    fmt.Println(fi6, p1, v36)
                                                    if len(v8) > 0 {
                                                        v8[0] = p1
                                                    }
                                                } else {
                                                    v35 -= v35
                                                    if true {
                                                        p0 = (p0 + (p0 / ((-9223372036854775808) | 1)))
                                                        v77 := GenS14{F0: 18446744073709551615, F1: -1}
                                                        v36 = 0.1
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            switch -406 {
                                                            case -2147483648:
                                                                p1 = p1
                                                            default:
                                                                v78 := GenS14{F0: 0, F1: (fi11 - -239)}
                                                                _ = v78
                                                            }
                                                        case "a":
                                                            v35 = (v36 / (1.5 * v35))
                                                            var v79 uint16 = 42
                                                            _ = v79
                                                        case "":
                                                            var v80 uint16 = 565
                                                            _ = v80
                                                        }
                                                        _ = len(v8)
                                                        if len(v8) > 0 {
                                                            v8[0] = (p1 ^ fi6)
                                                        }
                                                        if len(v8) > 0 {
                                                            v8[0] = 528
                                                        }
                                                        _ = v77
                                                    }
                                                }
                                                var v81 int16 = 1000
                                                _ = v81
                                            } else {
                                                _ = len(v8)
                                                switch 128 {
                                                case -563:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v82 uint8 = 2
                                                        _ = v82
                                                    case "\n\t":
                                                        switch 144 {
                                                        case -686:
                                                            if true {
                                                                v36 += (v36 * v36)
                                                                var v83 uint = 4294967295
                                                                fmt.Println(v36, v36, p0)
                                                                _ = v83
                                                            }
                                                        default:
                                                            v35 = float64(0.1)
                                                        }
                                                    case "a":
                                                        var v84 int = (p1 | 9223372036854775807)
                                                        _ = v84
                                                    default:
                                                        v8 = append(v8, p1)
                                                    }
                                                case -2147483648:
                                                    _ = len(v8)
                                                default:
                                                    var v85 int64 = p0
                                                    _ = v85
                                                }
                                                var v86 int16 = 42
                                                _ = v86
                                            }
                                            p0 = p0
                                            _ = v76
                                        case "\n\t":
                                            var v87 int8 = 123
                                            _ = v87
                                        }
                                        var v88 int64 = (p0 - int64(fi6))
                                        var v89 uint = 764
                                        _ = v88
                                        _ = v89
                                    }
                                    switch (fi6 / (((p1 >> ((37) & 63))) | 1)) {
                                    case 99:
                                        v35 -= v36
                                        v90 := GenS15{F0: fi11}
                                        var v91 int64 = -983
                                        _ = v90
                                        _ = v91
                                    case 128:
                                        switch (fi11 + (fi6 / (((fi11 << ((42) & 63))) | 1))) {
                                        case 1:
                                            var v92 uint = 2
                                            fmt.Println(fi11)
                                            var v93 bool = (2 > 2)
                                            var v94 float32 = float32((float64((1.5 - (3.14159 / 1.0))) - 0.0))
                                            _ = v92
                                            _ = v93
                                            _ = v94
                                        }
                                    case -380:
                                        switch 0 {
                                        case -2147483648:
                                            var v95 uint16 = 32767
                                            _ = v95
                                        case 277:
                                            if false {
                                                var v96 uint8 = 128
                                                p0 = (-(p0))
                                                _ = v96
                                            }
                                        }
                                    default:
                                        if ((false || true) || (false && (127 == 255))) {
                                            fmt.Println(p1, p1)
                                            if (!(0 == 0)) {
                                                var v97 int32 = 2147483647
                                                _ = v97
                                            }
                                            if true {
                                                fmt.Println(p1)
                                                if false {
                                                    fmt.Println(v36)
                                                } else {
                                                    var v98 float64 = v35
                                                    var v99 uint32 = 2
                                                    p1 -= (fi11 << ((2) & 63))
                                                    fmt.Println(fi6)
                                                    _ = v98
                                                    _ = v99
                                                }
                                            }
                                        } else {
                                            if (!(!true)) {
                                                var v100 uint64 = 4294967295
                                                _ = len(v8)
                                                if true {
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                }
                                                _ = v100
                                            } else {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v101 float32 = -2.25
                                                    _ = v101
                                                case "":
                                                    switch (fi6 % ((p1) | 1)) {
                                                    case -2147483648:
                                                        p1 -= 2147483647
                                                    case -1:
                                                        fmt.Println(fi6)
                                                    case -233:
                                                        fmt.Println(p0, fi11, p1)
                                                    default:
                                                        var v102 float32 = 0.1
                                                        _ = v102
                                                    }
                                                }
                                            }
                                            switch (p1 % (((fi11 / (((-(fi6))) | 1))) | 1)) {
                                            case -9223372036854775808:
                                                if false {
                                                    p1 = (fi6 << ((4294967295) & 63))
                                                } else {
                                                    p1 = -2147483648
                                                }
                                            }
                                            p0 += (p0 >> ((663) & 63))
                                        }
                                    }
                                }
                                switch 2147483647 {
                                case 42:
                                    fmt.Println(fi11)
                                }
                                _ = v35
                                _ = v36
                            }
                            var v103 int32 = 2147483647
                            if (65535 >= 946) {
                                fmt.Println(fi6, p1)
                            } else {
                                var v104 int16 = -809
                                fmt.Println(fi11)
                                _ = v104
                            }
                            v105 := GenS11{F0: fi11, F1: 2}
                            _ = v103
                            _ = v105
                        }
                    }
                    _ = fi11
                }
            default:
                v106 := GenS1{F0: 255, F1: -1}
                _ = v106
            }
        case "hello":
            var v107 uint64 = 2
            _ = v107
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 = 9223372036854775807
    v108 := GenS13{F0: -1, F1: 2147483648, F2: 42}
    v109 := GenS16{F0: 749}
    _ = p0
    _ = p1
    _ = v108
    _ = v109
    return (100 == -105)
}
func genFn1(p0 int, p1 int) uint64 {
    fmt.Println(p0, p0)
    for fi110 := 0; fi110 < 4; fi110++ {
        switch p1 {
        case 513:
            switch (-(p0)) {
            case 32768:
                if false {
                    fmt.Println(p0, p0)
                    for fi111 := 0; fi111 < 1; fi111++ {
                        p0 = (fi111 - -2147483648)
                        if (1 < 0) {
                            v112 := []int{}
                            switch "a" {
                            case "a":
                                switch 0 {
                                case 42:
                                    v113 := GenS12{F0: 9223372036854775807, F1: ((0.1 != 0.0) && (fi111 != (p1 ^ (p0 - -81))))}
                                    _ = v113
                                case 32767:
                                    v112 = append(v112, fi111)
                                case -1:
                                    if len(v112) >= 2 {
                                        v112 = v112[1:len(v112)]
                                    }
                                default:
                                    if len(v112) >= 2 {
                                        v112 = v112[1:len(v112)]
                                    }
                                }
                            case "":
                                if (1 <= 127) {
                                    if true {
                                        p0 = 1
                                        fmt.Println(p0)
                                    } else {
                                        v114 := GenS17{F0: 0, F1: 1, F2: 4294967295}
                                        if (81 < 4294967295) {
                                            if true {
                                                p0 -= (fi111 - (p1 * 4294967295))
                                                p1 += 9223372036854775807
                                            } else {
                                                p0 = (p0 + fi111)
                                            }
                                            fmt.Println(p1, p0)
                                            v115 := GenS14{F0: 4294967295, F1: (p0 * p1)}
                                            _ = v115
                                        }
                                        _ = v114
                                    }
                                    if (65535 == 185) {
                                        var v116 int = p1
                                        _ = v116
                                    }
                                    var v117 float64 = float64(1.5)
                                    _ = v117
                                } else {
                                    if len(v112) >= 2 {
                                        v112 = v112[1:len(v112)]
                                    }
                                }
                            case "hello":
                                switch p0 {
                                case 9223372036854775807:
                                    fmt.Println(p0, p0)
                                    if false {
                                        p0 += p0
                                        var v118 float32 = 1.5
                                        _ = v118
                                    }
                                case -82:
                                    if (42 != -9223372036854775808) {
                                        switch "hello" {
                                        case "\n\t":
                                            if ((!false) || (fi111 == p1)) {
                                                switch (("hello" + ("\n\t" + "\n\t")) + ("\n\t" + "foo bar")) {
                                                case "a":
                                                    switch (fi110 | (^(fi110))) {
                                                    case 42:
                                                        fmt.Println(fi110, p1)
                                                    case 9223372036854775807:
                                                        p0 -= p0
                                                    case 128:
                                                        switch fi111 {
                                                        case -362:
                                                            if (9223372036854775807 != 9223372036854775807) {
                                                                if (false && (429 < 255)) {
                                                                    fmt.Println(fi111, p0)
                                                                    p0 = 128
                                                                }
                                                                if false {
                                                                    v119 := GenS10{F0: -917}
                                                                    _ = v119
                                                                }
                                                                var v120 uint = 255
                                                                _ = v120
                                                            }
                                                        case 9223372036854775807:
                                                            v121 := GenS8{F0: 905, F1: 255, F2: 42}
                                                            _ = v121
                                                        case 32767:
                                                            var v122 float64 = (0.0 - float64(0.0))
                                                            _ = v122
                                                        }
                                                        var v123 uint8 = 255
                                                        _ = v123
                                                    }
                                                case "":
                                                    if false {
                                                        var v124 uint32 = 0
                                                        if len(v112) >= 2 {
                                                            v112 = v112[1:len(v112)]
                                                        }
                                                        var v125 uint32 = v124
                                                        _ = v124
                                                        _ = v125
                                                    }
                                                default:
                                                    var v126 bool = false
                                                    _ = v126
                                                }
                                            }
                                        case "hello":
                                            v127 := GenS2{F0: -27}
                                            _ = v127
                                        case "":
                                            v112 = append(v112, p0)
                                        }
                                        if (672 >= 466) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v128 := GenS2{F0: 127}
                                                _ = v128
                                            case "":
                                                if (!(-709 <= 359)) {
                                                    var v129 int = fi111
                                                    _ = v129
                                                }
                                            default:
                                                fmt.Println(fi111)
                                            }
                                            var v130 int8 = 127
                                            _ = v130
                                        }
                                    } else {
                                        fmt.Println(p1, fi111, fi110)
                                        p0 = -810
                                    }
                                case 1:
                                    if len(v112) > 0 {
                                        v112[0] = 127
                                    }
                                default:
                                    switch -2147483648 {
                                    case 4294967295:
                                        fmt.Println(fi111, fi111, fi111)
                                        v112 = append(v112, 615)
                                    }
                                }
                            default:
                                if (0 != 971) {
                                    p0 += p0
                                } else {
                                    var v131 uint8 = 68
                                    if false {
                                        p1 = fi111
                                        fmt.Println(fi110)
                                        if true {
                                            p1 += p1
                                            v132 := GenS18{F0: 901, F1: 2147483647}
                                            _ = v132
                                        }
                                    } else {
                                        if false {
                                            if false {
                                                var v133 float32 = float32(0.1)
                                                _ = v133
                                            } else {
                                                v131 = (v131 >> ((1) & 63))
                                                v134 := GenS19{F0: 752}
                                                p1 -= p0
                                                v135 := GenS20{F0: 1, F1: (65535 > 0)}
                                                fmt.Println(v131, fi110, p1)
                                                _ = v134
                                                _ = v135
                                            }
                                            var v136 float32 = 1e10
                                            _ = v136
                                        } else {
                                            var v137 int8 = 42
                                            _ = v137
                                        }
                                        p0 = p1
                                        var v138 int = (p1 - int(v131))
                                        _ = v138
                                    }
                                    v139 := GenS14{F0: 18446744073709551615, F1: p0}
                                    v131 = (^(v131))
                                    switch (^(fi110)) {
                                    case -923:
                                        if (!false) {
                                            switch "a" {
                                            case "a":
                                                var v140 uint16 = 0
                                                _ = v140
                                            case "foo bar":
                                                v141 := GenS13{F0: -1, F1: 65535, F2: 147}
                                                _ = v141
                                            }
                                            p0 -= p1
                                        }
                                    }
                                    _ = v131
                                    _ = v139
                                }
                            }
                            fmt.Println(fi111, fi111, fi110)
                            _ = v112
                        } else {
                            p0 += fi111
                        }
                        if (1 >= 174) {
                            v142 := GenS1{F0: 65535, F1: 42}
                            _ = v142
                        } else {
                            switch (("foo bar" + "hello") + ((("\n\t" + "a") + "a") + "\n\t")) {
                            case "a":
                                if (42 <= 0) {
                                    p1 = fi110
                                    switch "foo bar" {
                                    case "a":
                                        p1 -= 4294967295
                                    case "":
                                        p1 = (fi111 - fi111)
                                    case "\n\t":
                                        v143 := []int{}
                                        _ = v143
                                    }
                                    p0 -= fi110
                                }
                            case "hello":
                                fmt.Println(fi110, fi110, p0)
                            case "":
                                var v144 float32 = float32(-2.25)
                                _ = v144
                            default:
                                p1 = (fi110 >> ((795) & 63))
                            }
                            switch "foo bar" {
                            case "":
                                var v145 uint = 889
                                _ = v145
                            case "hello":
                                v146 := []int{(fi111 ^ p1), (p1 * (p1 & (p0 | (fi111 / ((fi111) | 1))))), p1, (fi110 | (p0 & 32768))}
                                _ = v146
                            case "\n\t":
                                if (!(1 == 1000)) {
                                    if false {
                                        p0 -= 32767
                                        switch 32768 {
                                        case -9223372036854775808:
                                            var v147 int64 = 9223372036854775807
                                            _ = v147
                                        case 128:
                                            v148 := []int{(-(p0)), (fi111 & (^(p0))), (fi110 + (^(fi111))), p0}
                                            _ = v148
                                        default:
                                            switch "a" {
                                            case "\n\t":
                                                var v149 int8 = 100
                                                _ = v149
                                            case "a":
                                                p1 -= p0
                                            case "foo bar":
                                                if true {
                                                    if true {
                                                        if (-17 == 127) {
                                                            v150 := GenS21{F0: 1, F1: -9223372036854775808, F2: 1}
                                                            v151 := GenS22{F0: true, F1: 128}
                                                            _ = v150
                                                            _ = v151
                                                        } else {
                                                            v152 := []int{-2147483648, (fi111 / (((fi110 - 32767)) | 1)), fi111}
                                                            p1 += (fi111 + fi110)
                                                            _ = v152
                                                        }
                                                    } else {
                                                        if (!(1 != 0)) {
                                                            var v153 int16 = 525
                                                            var v154 int8 = 127
                                                            _ = v153
                                                            _ = v154
                                                        } else {
                                                            var v155 string = ("\n\t" + "")
                                                            var v156 int = 1
                                                            var v157 string = (v155 + "")
                                                            _ = v155
                                                            _ = v156
                                                            _ = v157
                                                        }
                                                        v158 := GenS11{F0: 4294967295, F1: 0}
                                                        var v159 bool = (42 > 2)
                                                        if false {
                                                            var v160 int32 = 311
                                                            v161 := GenS2{F0: 42}
                                                            _ = v160
                                                            _ = v161
                                                        }
                                                        fmt.Println(fi111, v159, v159)
                                                        v162 := GenS2{F0: -128}
                                                        var v163 float32 = (1e10 * float32(0.0))
                                                        _ = v158
                                                        _ = v159
                                                        _ = v162
                                                        _ = v163
                                                    }
                                                    var v164 bool = true
                                                    _ = v164
                                                }
                                            default:
                                                fmt.Println(fi110)
                                            }
                                        }
                                        v165 := GenS5{F0: 2147483647}
                                        _ = v165
                                    } else {
                                        if true {
                                            if ((!false) || false) {
                                                var v166 string = "foo bar"
                                                fmt.Println(fi111, fi110)
                                                v167 := []int{}
                                                _ = v166
                                                _ = v167
                                            }
                                        }
                                        v168 := []int{(p0 - fi111)}
                                        v169 := GenS23{F0: 2}
                                        _ = v168
                                        _ = v169
                                    }
                                    switch ("" + "a") {
                                    case "\n\t":
                                        p1 -= -2147483648
                                    default:
                                        p1 = fi110
                                    }
                                    var v170 float64 = (1.5 * (0.1 - 3.14159))
                                    _ = v170
                                } else {
                                    if (false && false) {
                                        var v171 uint16 = 65535
                                        v172 := GenS24{F0: 128, F1: 14}
                                        var v173 int32 = 724
                                        _ = v171
                                        _ = v172
                                        _ = v173
                                    } else {
                                        var v174 int = fi111
                                        _ = v174
                                    }
                                    v175 := []int{(p1 & fi110), fi110, 32767, (^(fi111))}
                                    _ = v175
                                }
                            }
                            var v176 float64 = 3.14159
                            _ = v176
                        }
                        _ = fi111
                    }
                    v177 := []int{0, (p1 % (((^(p0))) | 1)), p1}
                    _ = v177
                } else {
                    v178 := []int{959, (p1 * fi110), fi110, fi110}
                    var v179 int32 = 0
                    _ = v178
                    _ = v179
                }
            case 32767:
                p0 = p0
            default:
                var v180 int = fi110
                _ = v180
            }
        case 2147483647:
            p1 = 32767
        case 218:
            p0 += (p0 ^ (fi110 - (p0 * (-(p1)))))
        default:
            v181 := []int{fi110, (^(fi110))}
            _ = v181
        }
        p0 = (p1 >> ((2) & 63))
        _ = fi110
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int = 128
    v2 := make(map[string]string)
    var v3 uint64 = 1
    var v4 uint64 = (^(v3))
    var v5 int = v1
    v1 += 25
    for fi182 := 0; fi182 < 15; fi182++ {
        switch "" {
        case "hello":
            if true {
                if ((v5 - 32768) >= fi182) {
                    var v183 int16 = -32768
                    var v184 uint8 = 128
                    _ = v183
                    _ = v184
                }
                var v185 uint32 = 65535
                v185 += 4294967295
                _ = v185
            } else {
                v2["\n\t"] = ""
                var v186 int32 = -1
                v2["foo bar"] = (("" + "a") + "hello")
                _ = v186
            }
        case "foo bar":
            for fi187 := 0; fi187 < 11; fi187++ {
                switch "" {
                case "hello":
                    var v188 float32 = (((1e10 / 0.1) - 0.1) * 1.5)
                    _ = v188
                case "a":
                    delete(v2, "a")
                    v4 = genFn1(fi182, v5)
                    if (v4 == genFn1(fi182, 0)) {
                        v189 := GenS2{F0: -128}
                        v2["a"] = "\n\t"
                        if (2147483648 != 2147483647) {
                            switch (-(v1)) {
                            case 42:
                                fmt.Println(fi187, v3)
                            case 32767:
                                if false {
                                    v4 = v3
                                    switch "" {
                                    case "hello":
                                        v190 := GenS23{F0: uint8(v3)}
                                        _ = v190
                                    default:
                                        v1 -= fi182
                                    }
                                    fmt.Println(v4, fi187, fi187)
                                    var v191 uint8 = 127
                                    var v192 uint16 = 858
                                    _ = v191
                                    _ = v192
                                }
                            case 32768:
                                v1 = v1
                                v3 += (v4 - (v4 % ((v3) | 1)))
                            }
                        }
                        _ = len(v2)
                        v193 := []int{-9223372036854775808, (v5 ^ fi182)}
                        v4 = (v4 - (v4 - 0))
                        v1 -= 128
                        var v194 int64 = -1
                        _ = v189
                        _ = v193
                        _ = v194
                    }
                default:
                    var v195 float32 = (3.14159 * (-2.25 - ((0.1 + 3.14159) / (3.14159 - float32((3.14159 * (0.1 + float64((0.1 * (-2.25 / (float32(0.1) * 1e10)))))))))))
                    _ = v195
                }
                fmt.Println(v1, v3, v5)
                v196 := GenS23{F0: 75}
                _ = len(v2)
                v5 = (fi187 * (fi182 | 972))
                v197 := GenS11{F0: (v5 / ((v1) | 1)), F1: (v4 % ((genFn1(610, (fi182 ^ (v1 * fi187)))) | 1))}
                v1 = v1
                _ = fi187
                _ = v196
                _ = v197
            }
        case "a":
            v3 -= (v4 & (v4 & v3))
        }
        _ = fi182
    }
    var v198 uint16 = 946
    v199 := GenS25{F0: ("a" + ""), F1: v5, F2: 111}
    for fi200 := 0; fi200 < 7; fi200++ {
        for fi201 := 0; fi201 < 19; fi201++ {
            var v202 int32 = 0
            var v203 uint32 = 791
            _ = fi201
            _ = v202
            _ = v203
        }
        _ = fi200
    }
    for fi204 := 0; fi204 < 15; fi204++ {
        delete(v2, "foo bar")
        fmt.Println(v3)
        v205 := GenS2{F0: -1}
        _ = fi204
        _ = v205
    }
    var v206 int = (v1 & (-(v5)))
    switch (v206 * 1) {
    case 1:
        switch "\n\t" {
        case "":
            v3 -= (v3 % (((v4 & (v3 << ((18446744073709551615) & 63)))) | 1))
        default:
            if true {
                v206 += -693
                var v207 int64 = -975
                var v208 float64 = 0.0
                _ = v207
                _ = v208
            } else {
                _ = len(v2)
            }
        }
    case 368:
        switch -1 {
        case 128:
            v209 := GenS6{F0: 42, F1: genFn1(v5, (v5 + (v1 % ((-2147483648) | 1)))), F2: (v206 / ((42) | 1))}
            _ = v209
        default:
            v5 = 9223372036854775807
        }
    case 32768:
        fmt.Println(v5, v198, v4)
    }
    if (!false) {
        fmt.Println(v4, v3)
        v210 := GenS25{F0: (("" + (("hello" + "") + "foo bar")) + ("\n\t" + ("hello" + "\n\t"))), F1: 127, F2: 904}
        _ = v210
    }
    var v211 uint32 = 42
    v1 = v5
    var v212 uint = 527
    fmt.Println("v1:", v1)
    fmt.Println("v198:", v198)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v206:", v206)
    fmt.Println("v211:", v211)
    fmt.Println("v212:", v212)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v198
    _ = v199
    _ = v2
    _ = v206
    _ = v211
    _ = v212
    _ = v3
    _ = v4
    _ = v5
}
