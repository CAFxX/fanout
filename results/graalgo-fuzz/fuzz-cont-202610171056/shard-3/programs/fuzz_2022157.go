package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int32; F1 string; F2 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8 }

func main() {
    v1 := []uint32{2147483647, 619, 659, 0}
    v2 := make(map[int]int)
    var v3 int32 = -1
    var v4 int16 = int16(v3)
    v4 -= v4
    switch ("a" + "a") {
    case "hello":
        var v5 float32 = 0.1
        _ = v5
    case "a":
        var v6 uint64 = 18446744073709551615
        _ = v6
    }
    v7 := GenS1{F0: 797}
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int = 42
        _ = fi8
        _ = v9
    }
    if true {
        v3 = (v3 & -1)
        fmt.Println(v4, v4)
        v3 = (^(v3))
    } else {
        var v10 int16 = (v4 - v4)
        var v11 int16 = (-(v10))
        fmt.Println(v3, v3)
        _ = v10
        _ = v11
    }
    switch "\n\t" {
    case "":
        fmt.Println(v4, v3)
    }
    v12 := GenS1{F0: -9223372036854775808}
    v4 = v4
    v13 := GenS2{F0: 1, F1: 128, F2: 168}
    if (false && (!(true && true))) {
        var v14 uint16 = 545
        _ = v14
    }
    v4 += (v4 ^ -924)
    if false {
        switch 4294967295 {
        case 42:
            v3 = -1
        case -9223372036854775808:
            switch 1 {
            case 4294967295:
                delete(v2, 1)
            default:
                var v15 uint32 = 20
                _ = v15
            }
        case 32767:
            switch -1 {
            case 42:
                v3 = (^(v3))
            case 2147483647:
                switch -9223372036854775808 {
                case 0:
                    v16 := GenS2{F0: 871, F1: 1, F2: 1}
                    _ = v16
                case -22:
                    v4 = v4
                default:
                    for fi17 := 0; fi17 < 4; fi17++ {
                        if true {
                            v3 -= (-(v3))
                        } else {
                            for fi18 := 0; fi18 < 6; fi18++ {
                                if true {
                                    v19 := GenS2{F0: 18446744073709551615, F1: 2, F2: 380}
                                    v4 = (v4 % ((v4) | 1))
                                    _ = v19
                                }
                                var v20 int8 = -102
                                _ = fi18
                                _ = v20
                            }
                        }
                        v3 += v3
                        _ = fi17
                    }
                }
            }
        }
        v4 = v4
        var v21 uint8 = 127
        _ = v21
    } else {
        switch 32768 {
        case 468:
            if (255 != 0) {
                if false {
                    var v22 uint16 = 1
                    _ = v22
                }
                if false {
                    fmt.Println(v4, v4)
                } else {
                    if (v4 != v4) {
                        var v23 int64 = 1
                        var v24 int64 = (v23 + 215)
                        _ = v23
                        _ = v24
                    } else {
                        var v25 int64 = 9223372036854775807
                        v3 = v3
                        fmt.Println(v3, v3, v25)
                        _ = v25
                    }
                    var v26 uint = 1
                    var v27 uint16 = uint16(v26)
                    _ = v26
                    _ = v27
                }
            } else {
                var v28 int32 = (-(v3))
                if (342 < 32767) {
                    if (42 == 15) {
                        if true {
                            for fi29 := 0; fi29 < 20; fi29++ {
                                switch fi29 {
                                case 1:
                                    for fi30 := 0; fi30 < 6; fi30++ {
                                        fmt.Println(fi29, v3)
                                        _ = fi30
                                    }
                                }
                                _ = fi29
                            }
                            for fi31 := 0; fi31 < 4; fi31++ {
                                var v32 int64 = 9223372036854775807
                                var v33 uint64 = 970
                                var v34 uint16 = 255
                                _ = fi31
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                        }
                        if false {
                            if (true || (true || false)) {
                                _ = len(v2)
                                v35 := GenS3{F0: -128}
                                _ = v35
                            } else {
                                v36 := GenS4{F0: 42}
                                _ = v36
                            }
                            if (!(true && ((0 > 255) && false))) {
                                if true {
                                    for fi37 := 0; fi37 < 18; fi37++ {
                                        if true {
                                            var v38 float32 = 1.5
                                            _ = v38
                                        }
                                        var v39 float32 = (-2.25 / 1.0)
                                        _ = fi37
                                        _ = v39
                                    }
                                    if ((v28 << ((1) & 63)) < 2147483647) {
                                        fmt.Println(v28, v3)
                                        fmt.Println(v3, v3)
                                        v1 = append(v1, 42)
                                    }
                                }
                                v40 := GenS5{F0: 0, F1: ((3.14159 == -2.25) && ((^(v4)) <= (v4 ^ (v4 | v4))))}
                                _ = v40
                            } else {
                                switch "a" {
                                case "a":
                                    for fi41 := 0; fi41 < 11; fi41++ {
                                        if (18446744073709551615 != 843) {
                                            v3 -= v3
                                            var v42 string = "a"
                                            _ = v42
                                        }
                                        fmt.Println(fi41)
                                        var v43 int64 = 9223372036854775807
                                        _ = fi41
                                        _ = v43
                                    }
                                case "hello":
                                    var v44 uint32 = 968
                                    _ = v44
                                default:
                                    var v45 uint32 = 738
                                    _ = v45
                                }
                                for fi46 := 0; fi46 < 18; fi46++ {
                                    if true {
                                        v28 += (v28 - (v28 + (^(v28))))
                                        v47 := v2[fi46]
                                        _ = v47
                                    } else {
                                        for fi48 := 0; fi48 < 18; fi48++ {
                                            fmt.Println(fi48, fi48)
                                            _ = fi48
                                        }
                                        if (354 != 18446744073709551615) {
                                            if (v28 > 42) {
                                                fmt.Println(fi46)
                                                var v49 string = ("hello" + "a")
                                                _ = v49
                                            } else {
                                                _ = len(v2)
                                                switch fi46 {
                                                case 1:
                                                    v50 := v2[fi46]
                                                    _ = v50
                                                case 4294967295:
                                                    if (-1 <= 2147483647) {
                                                        var v51 float32 = (1e10 - (-2.25 / 3.14159))
                                                        _ = v51
                                                    } else {
                                                        var v52 string = "foo bar"
                                                        var v53 uint64 = 28
                                                        _ = len(v2)
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                case 42:
                                                    var v54 uint16 = 65535
                                                    _ = v54
                                                }
                                            }
                                            v3 = v3
                                        } else {
                                            v55 := GenS1{F0: 762}
                                            v2[fi46] = 32768
                                            v56 := GenS3{F0: 0}
                                            _ = v55
                                            _ = v56
                                        }
                                        v4 = 647
                                    }
                                    _ = fi46
                                }
                                v57 := GenS4{F0: 1}
                                _ = v57
                            }
                            if ((!true) && true) {
                                switch 42 {
                                case -410:
                                    for fi58 := 0; fi58 < 9; fi58++ {
                                        var v59 int32 = v3
                                        _ = fi58
                                        _ = v59
                                    }
                                case -344:
                                    for fi60 := 0; fi60 < 19; fi60++ {
                                        var v61 int32 = 102
                                        v3 += v28
                                        var v62 int64 = -1
                                        _ = fi60
                                        _ = v61
                                        _ = v62
                                    }
                                }
                                v63 := GenS2{F0: 1, F1: 105, F2: 880}
                                _ = v63
                            }
                        } else {
                            v2[32768] = -873
                            v64 := GenS5{F0: 2147483647, F1: (((((-2.25 * -2.25) + 0.1) - (-2.25 / 1e10)) * 0.1) > (-2.25 - 0.1))}
                            var v65 bool = (!(true || ((false && (!false)) || false)))
                            _ = v64
                            _ = v65
                        }
                        if len(v1) > 0 {
                            v1[0] = 2147483647
                        }
                    }
                    var v66 uint32 = 819
                    _ = v66
                }
                if (!(1 >= 0)) {
                    for fi67 := 0; fi67 < 20; fi67++ {
                        for fi68 := 0; fi68 < 9; fi68++ {
                            v4 += (v4 | v4)
                            switch fi67 {
                            case -1:
                                v4 = 42
                            case 318:
                                if false {
                                    switch ((("" + "foo bar") + ("foo bar" + "")) + "") {
                                    case "a":
                                        if (!(!false)) {
                                            v69 := GenS2{F0: 18446744073709551615, F1: 21, F2: 0}
                                            _ = v69
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                var v70 int64 = 1
                                                _ = v70
                                            }
                                            fmt.Println(v3, v4)
                                        }
                                    case "hello":
                                        delete(v2, fi67)
                                    case "\n\t":
                                        v71 := GenS3{F0: 1}
                                        _ = v71
                                    }
                                } else {
                                    _ = len(v1)
                                    if (65535 != 1) {
                                        var v72 uint = 4294967295
                                        var v73 uint64 = 107
                                        _ = v72
                                        _ = v73
                                    } else {
                                        var v74 uint8 = 128
                                        var v75 int16 = (v4 & 1)
                                        _ = v74
                                        _ = v75
                                    }
                                    if true {
                                        if false {
                                            if (!false) {
                                                var v76 uint64 = 1
                                                var v77 float32 = 1e10
                                                var v78 bool = true
                                                _ = v76
                                                _ = v77
                                                _ = v78
                                            } else {
                                                _ = len(v1)
                                            }
                                            v79 := GenS1{F0: 0}
                                            if (false && false) {
                                                switch "foo bar" {
                                                case "":
                                                    fmt.Println(v3, v3)
                                                case "a":
                                                    if false {
                                                        switch fi67 {
                                                        case 9223372036854775807:
                                                            fmt.Println(v28)
                                                        case -2147483648:
                                                            fmt.Println(fi68, v3, v4)
                                                        default:
                                                            var v80 int16 = (v4 >> ((4294967295) & 63))
                                                            _ = v80
                                                        }
                                                        if (49 != 100) {
                                                            var v81 int64 = 1
                                                            v82 := GenS4{F0: 0}
                                                            v83 := GenS6{F0: (v28 << ((0) & 63)), F1: "hello", F2: (true && true)}
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                        v2[(fi68 | fi68)] = (-(fi68))
                                                        if (42 == 97) {
                                                            if false {
                                                                var v84 uint8 = 128
                                                                _ = v84
                                                            } else {
                                                                if (fi68 > fi68) {
                                                                    v85 := GenS7{F0: (false || (v4 == (v4 << ((0) & 63))))}
                                                                    _ = v85
                                                                }
                                                            }
                                                        }
                                                    }
                                                case "hello":
                                                    switch fi67 {
                                                    case 0:
                                                        v86 := GenS8{F0: -128}
                                                        _ = v86
                                                    case 4294967295:
                                                        if len(v1) > 0 {
                                                            v1[0] = 2
                                                        }
                                                    case 229:
                                                        v3 = v28
                                                    default:
                                                        var v87 float32 = (1.5 + 0.1)
                                                        _ = v87
                                                    }
                                                default:
                                                    switch "foo bar" {
                                                    case "a":
                                                        if (!((((-2.25 / -2.25) * float64((3.14159 + 1e10))) * 1e10) < (-2.25 * (1e10 / 1.0)))) {
                                                            v28 += (v28 >> ((42) & 63))
                                                            v3 += (v3 % ((1) | 1))
                                                            fmt.Println(fi67)
                                                        } else {
                                                            var v88 string = ""
                                                            if false {
                                                                var v89 int16 = (v4 / ((v4) | 1))
                                                                switch "" {
                                                                case "a":
                                                                    if true {
                                                                        if ((!false) && ((127 > 2) || false)) {
                                                                            var v90 int64 = 727
                                                                            var v91 int64 = (v90 + (-(v90)))
                                                                            switch (-(fi68)) {
                                                                            case 456:
                                                                                if (!false) {
                                                                                    var v92 int8 = 127
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 2147483647
                                                                                    }
                                                                                    _ = v92
                                                                                }
                                                                            case -625:
                                                                                if ((v88 + v88) >= v88) {
                                                                                    var v93 uint64 = 4294967295
                                                                                    if false {
                                                                                        v91 += v90
                                                                                        var v94 uint16 = 65535
                                                                                        _ = v94
                                                                                    }
                                                                                    var v95 uint = 445
                                                                                    _ = v93
                                                                                    _ = v95
                                                                                }
                                                                            default:
                                                                                fmt.Println(fi67, v4, v28)
                                                                            }
                                                                            _ = v90
                                                                            _ = v91
                                                                        }
                                                                    }
                                                                case "":
                                                                    _ = len(v2)
                                                                case "foo bar":
                                                                    fmt.Println(fi68)
                                                                default:
                                                                    v96 := v2[(fi68 * 1)]
                                                                    _ = v96
                                                                }
                                                                v4 = 1000
                                                                _ = v89
                                                            }
                                                            _ = v88
                                                        }
                                                    }
                                                }
                                            } else {
                                                var v97 uint32 = 2147483648
                                                switch fi68 {
                                                case -1:
                                                    var v98 bool = false
                                                    _ = v98
                                                case 4294967295:
                                                    if ((true && true) || false) {
                                                        var v99 uint64 = 1
                                                        switch (fi67 | (fi68 & (fi68 / ((fi68) | 1)))) {
                                                        case 32768:
                                                            var v100 bool = false
                                                            _ = v100
                                                        case 421:
                                                            v99 = v99
                                                        }
                                                        _ = v99
                                                    } else {
                                                        v3 = v3
                                                    }
                                                case 779:
                                                    _ = len(v1)
                                                default:
                                                    var v101 float64 = 3.14159
                                                    var v102 uint64 = 315
                                                    _ = v101
                                                    _ = v102
                                                }
                                                switch fi67 {
                                                case 158:
                                                    var v103 uint8 = 42
                                                    _ = v103
                                                case 128:
                                                    v28 += v3
                                                default:
                                                    v4 -= v4
                                                }
                                                _ = v97
                                            }
                                            _ = v79
                                        }
                                        if true {
                                            fmt.Println(v3, v3, fi68)
                                        } else {
                                            var v104 float32 = 0.0
                                            _ = v104
                                        }
                                    }
                                }
                            default:
                                v105 := v2[fi67]
                                _ = v105
                            }
                            v4 = 1
                            _ = fi68
                        }
                        _ = fi67
                    }
                    var v106 uint8 = 41
                    var v107 int8 = -1
                    _ = v106
                    _ = v107
                }
                _ = v28
            }
        case 127:
            _ = len(v2)
        case 1:
            var v108 bool = false
            _ = v108
        default:
            fmt.Println(v3, v3)
        }
        switch 596 {
        case 32767:
            fmt.Println(v4, v3)
        case 42:
            switch "a" {
            case "hello":
                _ = len(v1)
            default:
                v4 += v4
            }
        case 32768:
            v1 = append(v1, 65535)
        default:
            v4 = v4
        }
        if ((v4 % ((1000) | 1)) > v4) {
            fmt.Println(v3, v4)
        } else {
            v109 := GenS4{F0: 0}
            switch 9223372036854775807 {
            case 622:
                v110 := GenS8{F0: 127}
                _ = v110
            }
            for fi111 := 0; fi111 < 18; fi111++ {
                if (2 >= 1) {
                    switch ("\n\t" + "hello") {
                    case "":
                        v4 += (^(v4))
                    case "\n\t":
                        if false {
                            v4 = 42
                        } else {
                            var v112 bool = true
                            _ = v112
                        }
                    default:
                        for fi113 := 0; fi113 < 13; fi113++ {
                            switch "" {
                            case "a":
                                var v114 int64 = 9223372036854775807
                                _ = v114
                            case "hello":
                                v4 -= v4
                            default:
                                if true {
                                    v4 = 32767
                                    if (!((3.14159 * 0.0) == float64(((0.0 * 0.1) - (-2.25 / -2.25))))) {
                                        fmt.Println(v4, v3)
                                        fmt.Println(fi111, fi113)
                                        fmt.Println(v4, fi111, v3)
                                        var v115 uint = 18446744073709551615
                                        _ = v115
                                    }
                                    v3 += (v3 / (((v3 - (v3 - 426))) | 1))
                                } else {
                                    var v116 string = "foo bar"
                                    var v117 int8 = 1
                                    var v118 uint32 = 154
                                    _ = v116
                                    _ = v117
                                    _ = v118
                                }
                            }
                            if true {
                                delete(v2, (^(fi111)))
                            }
                            _ = fi113
                        }
                    }
                    v4 = (^(v4))
                }
                if true {
                    v119 := GenS8{F0: 100}
                    var v120 float32 = (-2.25 - (3.14159 * 1e10))
                    fmt.Println(v4)
                    _ = v119
                    _ = v120
                } else {
                    for fi121 := 0; fi121 < 13; fi121++ {
                        if ((v4 << ((2) & 63)) <= v4) {
                            v122 := GenS7{F0: (true || false)}
                            if false {
                                v3 += v3
                            } else {
                                var v123 int = (fi121 | fi121)
                                _ = v123
                            }
                            _ = v122
                        }
                        var v124 int32 = (v3 << ((4294967295) & 63))
                        _ = fi121
                        _ = v124
                    }
                    fmt.Println(v4)
                    var v125 int64 = 9223372036854775807
                    _ = v125
                }
                var v126 uint = 4294967295
                _ = fi111
                _ = v126
            }
            _ = v109
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
