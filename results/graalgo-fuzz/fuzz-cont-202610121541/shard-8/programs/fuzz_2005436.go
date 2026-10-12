package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 bool; F1 uint8; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int16; F2 int8 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint; F1 int; F2 string }
type GenS10 struct { F0 uint8; F1 int64 }
type GenS11 struct { F0 uint8; F1 bool }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int16; F1 uint; F2 uint64 }
type GenS14 struct { F0 uint; F1 uint16 }
type GenS15 struct { F0 uint; F1 uint64 }
type GenS16 struct { F0 int; F1 int32 }
type GenS17 struct { F0 string; F1 uint8; F2 uint32 }

func main() {
    v1 := []uint32{}
    var v2 uint = 1
    var v3 uint = (v2 >> ((v2) & 63))
    var v4 uint64 = uint64(v3)
    var v5 int = 1000
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v6 uint16 = 1
    var v7 uint64 = (^(v4))
    v2 = v2
    for fi8 := 0; fi8 < 6; fi8++ {
        switch ("hello" + "foo bar") {
        case "hello":
            fmt.Println(v7, fi8)
        case "a":
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(v7, fi9, v5)
                fmt.Println(v2)
                var v10 uint8 = 155
                _ = fi9
                _ = v10
            }
        case "foo bar":
            _ = len(v1)
        default:
            _ = len(v1)
        }
        v1 = append(v1, 900)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS1{F0: -128, F1: 4294967295}
        var v13 int = (fi11 + (v5 << ((v3) & 63)))
        var v14 int = (^(v13))
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    if (v5 != v5) {
        var v15 int32 = 0
        switch v5 {
        case -1:
            v16 := GenS1{F0: 42, F1: (v2 & (v2 / (((v3 % ((v3) | 1))) | 1)))}
            _ = v16
        }
        _ = v15
    } else {
        fmt.Println(v3, v5)
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        fmt.Println(v4, v7, v3)
        fmt.Println(v2, v5)
        if len(v1) > 0 {
            v1[0] = uint32(v3)
        }
        _ = fi17
    }
    v18 := GenS2{F0: true, F1: 255, F2: 42}
    if (217 == 2) {
        switch v5 {
        case -9223372036854775808:
            var v19 float64 = float64(0.1)
            _ = v19
        case -401:
            var v20 bool = (v6 >= 64)
            _ = v20
        default:
            if (((0.0 + 1e10) * 0.0) == 1e10) {
                var v21 int64 = 999
                v2 = 284
                _ = v21
            } else {
                var v22 string = "a"
                if ((!true) && (!(v6 < 318))) {
                    v3 += v2
                    v5 = (v5 % (((v5 + (v5 ^ (v5 | v5)))) | 1))
                    v5 += 128
                } else {
                    v5 -= -945
                }
                _ = v22
            }
        }
        if false {
            v7 += 421
            for fi23 := 0; fi23 < 7; fi23++ {
                for fi24 := 0; fi24 < 18; fi24++ {
                    if (true || (!(42 != 2147483647))) {
                        switch v5 {
                        case -97:
                            v4 = 402
                        case 32768:
                            var v25 uint64 = (v7 & 18446744073709551615)
                            v5 -= 127
                            v2 += v2
                            _ = v25
                        default:
                            fmt.Println(v3, fi24)
                        }
                        v6 = v6
                        var v26 int64 = 9223372036854775807
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v26
                    }
                    _ = fi24
                }
                var v27 float32 = (((((1e10 - 1e10) / -2.25) - (((-2.25 - 1.5) * 0.1) * (0.0 / 1.5))) - 3.14159) - -2.25)
                var v28 int16 = int16(v3)
                _ = fi23
                _ = v27
                _ = v28
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := GenS1{F0: 127, F1: uint(v6)}
        if false {
            v5 = (-(fi29))
            if true {
                for fi31 := 0; fi31 < 14; fi31++ {
                    var v32 float64 = ((0.0 + 1.5) * -2.25)
                    v32 += float64(((1e10 / 0.1) - (((3.14159 / 1.0) * ((0.1 / float32((float64(0.0) * ((v32 / 1.5) / 1e10)))) + float32((3.14159 - 0.0)))) * (1e10 + 1e10))))
                    _ = fi31
                    _ = v32
                }
                var v33 uint16 = v6
                _ = v33
            } else {
                for fi34 := 0; fi34 < 11; fi34++ {
                    if (fi29 < (fi29 & (-(v5)))) {
                        var v35 bool = true
                        var v36 uint8 = 255
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                        var v37 uint64 = v7
                        v4 -= v37
                        _ = v35
                        _ = v36
                        _ = v37
                    } else {
                        switch -885 {
                        case 1:
                            if true {
                                v4 -= (v4 * v7)
                                if len(v1) > 0 {
                                    v1[0] = 958
                                }
                            } else {
                                v38 := GenS2{F0: (860 == 765), F1: 1, F2: 0}
                                switch (fi29 % ((860) | 1)) {
                                case 332:
                                    v4 += v7
                                    v39 := GenS2{F0: true, F1: 145, F2: 1}
                                    v6 += 750
                                    _ = v39
                                case 42:
                                    switch 0 {
                                    case 81:
                                        fmt.Println(v2)
                                    case 128:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v40 uint8 = 42
                                        if (2147483648 >= uint32(v40)) {
                                            v41 := GenS3{F0: (v5 / ((fi29) | 1))}
                                            if false {
                                                switch (v5 * 528) {
                                                case 127:
                                                    v4 -= (v4 + v7)
                                                case 42:
                                                    fmt.Println(v4)
                                                case 4294967295:
                                                    var v42 int = (fi29 << (((v2 | v3)) & 63))
                                                    _ = v42
                                                }
                                                v43 := GenS3{F0: (v5 | fi34)}
                                                v44 := GenS2{F0: false, F1: v40, F2: int8(v6)}
                                                if len(v1) > 0 {
                                                    v1[0] = 315
                                                }
                                                _ = v43
                                                _ = v44
                                            }
                                            v6 = v6
                                            _ = v41
                                        } else {
                                            v2 = v2
                                            v45 := GenS3{F0: fi34}
                                            _ = v45
                                        }
                                        _ = v40
                                    }
                                default:
                                    v6 -= 32767
                                }
                                fmt.Println(v2, v2)
                                if false {
                                    switch fi29 {
                                    case 32768:
                                        var v46 uint64 = (v4 - v7)
                                        _ = v46
                                    case -2147483648:
                                        v47 := GenS2{F0: true, F1: 42, F2: 0}
                                        _ = v47
                                    }
                                } else {
                                    var v48 string = "a"
                                    if (!((-77 != 199) && false)) {
                                        if (!false) {
                                            var v49 int8 = 0
                                            _ = v49
                                        }
                                        v7 = 0
                                        if (v5 > fi29) {
                                            if (369 > 0) {
                                                v7 += (v7 & 18446744073709551615)
                                            }
                                            if (42 != -2147483648) {
                                                var v50 uint8 = 127
                                                _ = v50
                                            }
                                            v51 := GenS1{F0: 100, F1: v3}
                                            fmt.Println(fi29)
                                            _ = v51
                                        } else {
                                            if (true || false) {
                                                switch (fi34 << ((65535) & 63)) {
                                                case 0:
                                                    var v52 int8 = 1
                                                    _ = v52
                                                case 128:
                                                    fmt.Println(v2, v2, v4)
                                                    v2 = v3
                                                case 4294967295:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if false {
                                                            var v53 uint64 = (v7 & v4)
                                                            if ((!(true || false)) || false) {
                                                                var v54 bool = (true || false)
                                                                v1 = append(v1, 1)
                                                                _ = v54
                                                            } else {
                                                                var v55 uint = 2
                                                                fmt.Println(v2, v4, v3)
                                                                switch v5 {
                                                                case -404:
                                                                    v56 := GenS4{F0: 166, F1: -32768, F2: int8(v53)}
                                                                    _ = v56
                                                                case -1:
                                                                    v4 -= v53
                                                                case 127:
                                                                    v57 := GenS5{F0: int(v7), F1: (^(v53))}
                                                                    _ = v57
                                                                }
                                                                _ = v55
                                                            }
                                                            var v58 int64 = 884
                                                            _ = v53
                                                            _ = v58
                                                        } else {
                                                            var v59 uint16 = 32767
                                                            v60 := GenS6{F0: 1000, F1: 1}
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v61 uint16 = 65535
                                                                _ = v61
                                                            default:
                                                                var v62 int32 = -2147483648
                                                                _ = v62
                                                            }
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    }
                                                default:
                                                    v3 = (v3 - 320)
                                                    v5 = (v5 / (((fi29 >> (((v2 % (((v3 % ((v3) | 1))) | 1))) & 63))) | 1))
                                                    v2 = 547
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = len(v1)
                                                    v63 := GenS7{F0: 0, F1: (v4 % (((v4 >> (((v3 & (v2 & (v3 ^ v3)))) & 63))) | 1))}
                                                    _ = v63
                                                case "foo bar":
                                                    fmt.Println(v3, v48)
                                                case "":
                                                    var v64 float32 = 0.1
                                                    _ = v64
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                v65 := GenS1{F0: -21, F1: (v3 | v3)}
                                                _ = v65
                                            } else {
                                                switch 2147483647 {
                                                case 128:
                                                    switch v48 {
                                                    case "a":
                                                        switch "" {
                                                        case "":
                                                            v48 = v48
                                                            var v66 int8 = -128
                                                            v67 := GenS8{F0: -283}
                                                            _ = v66
                                                            _ = v67
                                                        case "hello":
                                                            if (-128 != 0) {
                                                                if len(v1) > 0 {
                                                                    v1[0] = 65535
                                                                }
                                                            } else {
                                                                if (!(-1 == -1)) {
                                                                    switch fi34 {
                                                                    case 0:
                                                                        v7 = v7
                                                                    default:
                                                                        var v68 int8 = int8(fi34)
                                                                        var v69 int16 = int16(v6)
                                                                        _ = v68
                                                                        _ = v69
                                                                    }
                                                                    var v70 uint32 = 2147483647
                                                                    _ = len(v1)
                                                                    _ = v70
                                                                }
                                                                switch "" {
                                                                case "":
                                                                    var v71 bool = true
                                                                    _ = v71
                                                                case "hello":
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    var v72 int64 = -443
                                                                    _ = v72
                                                                default:
                                                                    if (-43 != -34) {
                                                                        if true {
                                                                            if true {
                                                                                fmt.Println(v3, fi29)
                                                                            } else {
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                if true {
                                                                                    fmt.Println(v2, v6, v4)
                                                                                    switch v5 {
                                                                                    case 127:
                                                                                        v73 := GenS9{F0: v2, F1: fi29, F2: v48}
                                                                                        _ = v73
                                                                                    case 42:
                                                                                        var v74 uint8 = 127
                                                                                        _ = v74
                                                                                    case 2147483647:
                                                                                        var v75 bool = false
                                                                                        _ = v75
                                                                                    default:
                                                                                        var v76 uint = 707
                                                                                        _ = v76
                                                                                    }
                                                                                } else {
                                                                                    switch int(v6) {
                                                                                    case 25:
                                                                                        v77 := GenS4{F0: -9223372036854775808, F1: -32768, F2: -116}
                                                                                        if (false || ((v4 << (((v3 ^ v3)) & 63)) < 573)) {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = 937
                                                                                            }
                                                                                        } else {
                                                                                            var v78 int = (-(v5))
                                                                                            _ = v78
                                                                                        }
                                                                                        _ = v77
                                                                                    case 32768:
                                                                                        v3 = v3
                                                                                    }
                                                                                    var v79 int32 = 582
                                                                                    if (v4 > uint64(v6)) {
                                                                                        v7 += 42
                                                                                        var v80 int32 = 42
                                                                                        v79 = v80
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 2147483647
                                                                                        }
                                                                                        _ = v80
                                                                                    }
                                                                                    _ = v79
                                                                                }
                                                                            }
                                                                        }
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        var v81 bool = ((42 <= 127) || false)
                                                                        _ = v81
                                                                    } else {
                                                                        switch v48 {
                                                                        case "":
                                                                            fmt.Println(v6, v7, v4)
                                                                        case "a":
                                                                            switch (fi34 | fi29) {
                                                                            case 128:
                                                                                v82 := GenS1{F0: 0, F1: (v2 >> ((v2) & 63))}
                                                                                _ = v82
                                                                            }
                                                                            v6 += uint16(v4)
                                                                            if (!(0.1 < 1.5)) {
                                                                                var v83 uint8 = 2
                                                                                switch v48 {
                                                                                case "\n\t":
                                                                                    var v84 bool = (uint16(v4) < 1)
                                                                                    _ = v84
                                                                                default:
                                                                                    var v85 uint8 = v83
                                                                                    _ = v85
                                                                                }
                                                                                var v86 uint8 = (v83 * 0)
                                                                                _ = v83
                                                                                _ = v86
                                                                            }
                                                                            var v87 float64 = float64(-2.25)
                                                                            if false {
                                                                                switch v48 {
                                                                                case "\n\t":
                                                                                    v4 = v4
                                                                                case "":
                                                                                    v88 := GenS1{F0: -34, F1: (v2 - v3)}
                                                                                    _ = v88
                                                                                }
                                                                                if true {
                                                                                    if (true || ((255 != 127) || false)) {
                                                                                        v89 := GenS10{F0: 42, F1: 193}
                                                                                        _ = len(v1)
                                                                                        if true {
                                                                                            v90 := GenS2{F0: ((1e10 + 1e10) != (1.5 + float32((v87 - ((((float64(0.0) - v87) + 1e10) / v87) + (float64((float32((float64(3.14159) - v87)) * 1e10)) * v87)))))), F1: 0, F2: -1}
                                                                                            v5 = v5
                                                                                            _ = v90
                                                                                        }
                                                                                        _ = v89
                                                                                    }
                                                                                    v1 = append(v1, 2147483647)
                                                                                    v7 = v7
                                                                                } else {
                                                                                    var v91 uint = (v2 % ((v3) | 1))
                                                                                    v92 := GenS11{F0: 255, F1: (-532 == int64(v3))}
                                                                                    if (128 >= 255) {
                                                                                        _ = len(v1)
                                                                                    }
                                                                                    _ = v91
                                                                                    _ = v92
                                                                                }
                                                                            } else {
                                                                                if false {
                                                                                    if (false || (!(42 == 42))) {
                                                                                        v3 = v3
                                                                                    }
                                                                                } else {
                                                                                    v48 = v48
                                                                                    v6 = (^(v6))
                                                                                    if false {
                                                                                        var v93 uint16 = (v6 + 772)
                                                                                        if false {
                                                                                            var v94 bool = ((^(v4)) == v4)
                                                                                            _ = v94
                                                                                        } else {
                                                                                            var v95 float32 = (0.1 - (float32((v87 * v87)) - 0.1))
                                                                                            var v96 int8 = 0
                                                                                            _ = v95
                                                                                            _ = v96
                                                                                        }
                                                                                        _ = v93
                                                                                    }
                                                                                    var v97 uint8 = 127
                                                                                    _ = v97
                                                                                }
                                                                                var v98 int32 = 2147483647
                                                                                _ = v98
                                                                            }
                                                                            var v99 int8 = -102
                                                                            switch (^(fi34)) {
                                                                            case 990:
                                                                                var v100 uint32 = 2147483647
                                                                                _ = v100
                                                                            case 9223372036854775807:
                                                                                v101 := GenS2{F0: (-392 != -2147483648), F1: 9, F2: v99}
                                                                                _ = v101
                                                                            default:
                                                                                v102 := GenS3{F0: (v5 & v5)}
                                                                                _ = v102
                                                                            }
                                                                            _ = v87
                                                                            _ = v99
                                                                        case "\n\t":
                                                                            fmt.Println(v5)
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case "a":
                                                            switch v48 {
                                                            case "hello":
                                                                v103 := GenS12{F0: 42}
                                                                if true {
                                                                    var v104 uint8 = 230
                                                                    var v105 float64 = 1e10
                                                                    _ = v104
                                                                    _ = v105
                                                                } else {
                                                                    v1 = append(v1, 716)
                                                                    if true {
                                                                        fmt.Println(v6)
                                                                        var v106 uint64 = v7
                                                                        _ = v106
                                                                    } else {
                                                                        v7 -= 1
                                                                    }
                                                                    var v107 uint32 = 0
                                                                    _ = v107
                                                                }
                                                                _ = v103
                                                            case "foo bar":
                                                                v5 = -532
                                                            default:
                                                                var v108 int = int(v4)
                                                                _ = v108
                                                            }
                                                        default:
                                                            if true {
                                                                v2 -= v2
                                                                var v109 float64 = 3.14159
                                                                v2 = v3
                                                                v6 -= 255
                                                                _ = v109
                                                            }
                                                        }
                                                    case "foo bar":
                                                        v110 := GenS4{F0: -9223372036854775808, F1: -81, F2: -7}
                                                        v111 := GenS13{F0: 32767, F1: v2, F2: uint64(v3)}
                                                        _ = v110
                                                        _ = v111
                                                    default:
                                                        var v112 float32 = (float32((1.5 + 0.0)) - ((0.1 - 0.1) + 1e10))
                                                        _ = v112
                                                    }
                                                case -2147483648:
                                                    fmt.Println(fi34, v6, fi34)
                                                default:
                                                    switch v48 {
                                                    case "foo bar":
                                                        v7 += v7
                                                        if (2147483647 < 0) {
                                                            if true {
                                                                v48 = v48
                                                            }
                                                            v5 -= 746
                                                        } else {
                                                            var v113 string = "hello"
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v114 float64 = 1.5
                                                            v5 = 32768
                                                            _ = v113
                                                            _ = v114
                                                        }
                                                    case "hello":
                                                        v115 := GenS2{F0: true, F1: 60, F2: int8(v4)}
                                                        _ = v115
                                                    case "\n\t":
                                                        if false {
                                                            if (-128 < 100) {
                                                                if (9223372036854775807 < -1) {
                                                                    v3 += v2
                                                                    var v116 uint = uint(v6)
                                                                    _ = v116
                                                                } else {
                                                                    var v117 uint64 = (v7 & (v7 % (((v4 * 18446744073709551615)) | 1)))
                                                                    v1 = append(v1, 114)
                                                                    v117 -= v7
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 2147483648
                                                                    }
                                                                    _ = v117
                                                                }
                                                                switch "hello" {
                                                                case "":
                                                                    if false {
                                                                        if false {
                                                                            v7 = (v7 | 385)
                                                                            switch "hello" {
                                                                            case "a":
                                                                                fmt.Println(fi29, v3, v6)
                                                                            case "":
                                                                                var v118 float32 = ((-2.25 * 0.1) - (0.1 - 3.14159))
                                                                                _ = v118
                                                                            case "\n\t":
                                                                                var v119 int32 = -2147483648
                                                                                _ = v119
                                                                            default:
                                                                                var v120 int64 = 42
                                                                                _ = v120
                                                                            }
                                                                        }
                                                                        var v121 uint8 = 1
                                                                        var v122 bool = (!(false && false))
                                                                        _ = v121
                                                                        _ = v122
                                                                    } else {
                                                                        var v123 string = ((v48 + (v48 + "")) + v48)
                                                                        v124 := GenS14{F0: v3, F1: v6}
                                                                        v5 = v5
                                                                        var v125 string = ""
                                                                        _ = v123
                                                                        _ = v124
                                                                        _ = v125
                                                                    }
                                                                }
                                                            }
                                                            var v126 int64 = -9223372036854775808
                                                            var v127 int = (-(fi34))
                                                            if len(v1) > 0 {
                                                                v1[0] = 4294967295
                                                            }
                                                            v1 = append(v1, 42)
                                                            _ = v126
                                                            _ = v127
                                                        } else {
                                                            v1 = append(v1, 2)
                                                            var v128 string = ("\n\t" + "\n\t")
                                                            _ = v128
                                                        }
                                                    }
                                                }
                                                v2 = (^(v2))
                                                var v129 int32 = 0
                                                var v130 int8 = 89
                                                var v131 uint64 = v4
                                                _ = v129
                                                _ = v130
                                                _ = v131
                                            }
                                            switch 32768 {
                                            case 0:
                                                fmt.Println(v2, fi29)
                                            case 42:
                                                if true {
                                                    if (!false) {
                                                        v1 = append(v1, 0)
                                                    } else {
                                                        var v132 uint64 = (v7 | 816)
                                                        switch ("" + (v48 + v48)) {
                                                        case "\n\t":
                                                            if (false || (fi34 == (fi29 ^ fi34))) {
                                                                if true {
                                                                    var v133 bool = (!(true && false))
                                                                    _ = v133
                                                                } else {
                                                                    fmt.Println(v5, v48)
                                                                    v5 += v5
                                                                    var v134 uint64 = (^(v132))
                                                                    _ = v134
                                                                }
                                                            }
                                                        case "foo bar":
                                                            v7 = (v132 - (v7 - v7))
                                                        case "hello":
                                                            v7 = (^(v7))
                                                        }
                                                        _ = v132
                                                    }
                                                    v5 += v5
                                                    if (0.1 < (1e10 + -2.25)) {
                                                        v1 = append(v1, 0)
                                                    } else {
                                                        v135 := GenS13{F0: 1, F1: v3, F2: 726}
                                                        _ = v135
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = 4294967295
                                                    }
                                                }
                                            case -1:
                                                if true {
                                                    v136 := GenS15{F0: (^(v3)), F1: 18446744073709551615}
                                                    v2 = (v3 | v3)
                                                    _ = v136
                                                } else {
                                                    v1 = append(v1, 1)
                                                    var v137 int32 = 0
                                                    _ = v137
                                                }
                                            }
                                        }
                                    }
                                    switch (fi34 ^ 127) {
                                    case 0:
                                        v138 := GenS15{F0: (v2 & v3), F1: 2}
                                        _ = v138
                                    case -177:
                                        var v139 uint16 = (^(v6))
                                        _ = v139
                                    }
                                    _ = v48
                                }
                                _ = v38
                            }
                        case 905:
                            switch "" {
                            case "hello":
                                var v140 string = ("foo bar" + ("a" + "\n\t"))
                                _ = v140
                            case "":
                                var v141 float32 = (0.1 - 1.5)
                                _ = v141
                            }
                        default:
                            v142 := GenS6{F0: 32767, F1: v2}
                            _ = v142
                        }
                        v4 += 0
                    }
                    var v143 uint32 = 2147483648
                    if false {
                        switch (fi29 - int(v2)) {
                        case 938:
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                        case -1:
                            switch "\n\t" {
                            case "foo bar":
                                if (v4 == 2) {
                                    switch "hello" {
                                    case "\n\t":
                                        v7 += (v4 | (v4 - v7))
                                    case "hello":
                                        var v144 bool = (-120 <= 100)
                                        _ = v144
                                    case "foo bar":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                }
                            case "a":
                                switch 961 {
                                case 128:
                                    var v145 int16 = 1000
                                    _ = v145
                                case 798:
                                    var v146 uint8 = 42
                                    _ = v146
                                case -9223372036854775808:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            default:
                                var v147 int16 = 42
                                _ = v147
                            }
                        case 9223372036854775807:
                            var v148 int16 = -32768
                            _ = v148
                        }
                        v149 := GenS16{F0: (fi34 << ((v3) & 63)), F1: 42}
                        v150 := GenS17{F0: "\n\t", F1: 197, F2: v143}
                        _ = v149
                        _ = v150
                    }
                    _ = fi34
                    _ = v143
                }
            }
        }
        _ = fi29
        _ = v30
    }
    v3 -= 1
    var v151 int8 = 1
    var v152 uint = v3
    v4 = (v7 + (^(v7)))
    var v153 string = ("" + ("\n\t" + "a"))
    v154 := GenS17{F0: (((v153 + v153) + v153) + v153), F1: 0, F2: 65535}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v151
    _ = v152
    _ = v153
    _ = v154
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
