package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int64; F1 int64; F2 int8 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int16; F1 int64; F2 uint8 }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 uint16; F1 uint16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 bool; F1 string }
type GenS14 struct { F0 int8; F1 int64; F2 int8 }
type GenS15 struct { F0 uint64; F1 uint16 }
type GenS16 struct { F0 uint16; F1 uint }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 uint64; F1 int32; F2 int }
type GenS19 struct { F0 uint16; F1 int64; F2 int8 }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 int64 }

func main() {
    var v1 float32 = (-2.25 + (((3.14159 - 0.1) + 3.14159) / -2.25))
    v2 := []uint32{}
    var v3 string = "foo bar"
    v4 := []uint32{2, 42}
    var v5 int16 = -1
    v6 := []int{42, -1, 141, 32768}
    v3 = "\n\t"
    fmt.Println(v3, v5)
    if true {
        var v7 int32 = 1
        if ((!true) && (false && true)) {
            switch v3 {
            case "foo bar":
                v5 += 956
            case "hello":
                if (684 >= 284) {
                    v8 := GenS1{F0: 134}
                    switch (v3 + v3) {
                    case "":
                        fmt.Println(v3)
                    case "a":
                        if true {
                            v7 = (v7 & -813)
                        } else {
                            var v9 string = v3
                            var v10 int64 = 683
                            if len(v6) > 0 {
                                v6[0] = 2147483647
                            }
                            _ = v10
                            _ = v9
                        }
                    case "foo bar":
                        var v11 int = 2147483647
                        _ = v11
                    default:
                        if false {
                            for fi12 := 0; fi12 < 18; fi12++ {
                                for fi13 := 0; fi13 < 6; fi13++ {
                                    var v14 int8 = -1
                                    _ = fi13
                                    _ = v14
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v5 += (v5 >> ((437) & 63))
                                _ = fi12
                            }
                            v15 := GenS2{F0: 1, F1: -1}
                            if ((-(v5)) != v5) {
                                v16 := GenS2{F0: 128, F1: 73}
                                v3 = v3
                                _ = v16
                            }
                            _ = v15
                        } else {
                            fmt.Println(v5)
                            switch 9223372036854775807 {
                            case -617:
                                var v17 uint16 = 40
                                _ = v17
                            case 0:
                                for fi18 := 0; fi18 < 7; fi18++ {
                                    v19 := GenS3{F0: 18446744073709551615, F1: 18446744073709551615, F2: 42}
                                    var v20 uint = 481
                                    v21 := GenS2{F0: 1, F1: -1}
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                            case -714:
                                for fi22 := 0; fi22 < 3; fi22++ {
                                    fmt.Println(v7, v3, fi22)
                                    var v23 int = fi22
                                    _ = fi22
                                    _ = v23
                                }
                            }
                            if (42 != 4294967295) {
                                v1 = 1.5
                            } else {
                                if false {
                                    fmt.Println(v1, v7, v1)
                                }
                            }
                        }
                    }
                    _ = v8
                } else {
                    var v24 int64 = 9223372036854775807
                    if true {
                        switch v3 {
                        case "":
                            for fi25 := 0; fi25 < 1; fi25++ {
                                if len(v6) > 0 {
                                    v6[0] = fi25
                                }
                                switch (v3 + (v3 + (v3 + ((v3 + (v3 + (v3 + (v3 + ("foo bar" + "hello"))))) + "\n\t")))) {
                                case "hello":
                                    if true {
                                        var v26 int32 = -2147483648
                                        _ = v26
                                    } else {
                                        var v27 int64 = 0
                                        for fi28 := 0; fi28 < 16; fi28++ {
                                            if false {
                                                v1 = v1
                                                var v29 uint16 = 657
                                                var v30 bool = (!true)
                                                _ = v29
                                                _ = v30
                                            } else {
                                                v31 := GenS4{F0: (v24 - (v27 << ((557) & 63))), F1: (v24 * (v27 << ((18446744073709551615) & 63))), F2: -115}
                                                _ = v31
                                            }
                                            _ = fi28
                                        }
                                        if false {
                                            var v32 int16 = 1
                                            v27 += v24
                                            _ = v32
                                        }
                                        _ = v27
                                    }
                                case "a":
                                    for fi33 := 0; fi33 < 16; fi33++ {
                                        fmt.Println(v7)
                                        v34 := GenS5{F0: 42, F1: v24}
                                        v7 -= (-(v7))
                                        _ = fi33
                                        _ = v34
                                    }
                                default:
                                    switch (-(fi25)) {
                                    case 128:
                                        var v35 string = ""
                                        _ = v35
                                    case 32767:
                                        var v36 uint64 = 0
                                        _ = v36
                                    case 1:
                                        for fi37 := 0; fi37 < 19; fi37++ {
                                            fmt.Println(v7, v1, v7)
                                            _ = fi37
                                        }
                                    default:
                                        v4 = append(v4, 42)
                                    }
                                }
                                _ = fi25
                            }
                        case "hello":
                            var v38 int8 = -1
                            _ = v38
                        case "\n\t":
                            switch -2147483648 {
                            case -9223372036854775808:
                                switch "hello" {
                                case "a":
                                    v7 = -1
                                default:
                                    v2 = append(v2, 2147483648)
                                }
                            case 2147483647:
                                v2 = append(v2, 4294967295)
                            default:
                                var v39 uint16 = 4
                                _ = v39
                            }
                        default:
                            v40 := GenS1{F0: 106}
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 18; fi41++ {
                            fmt.Println(v7, v1)
                            _ = fi41
                        }
                        for fi42 := 0; fi42 < 16; fi42++ {
                            v43 := GenS6{F0: "hello"}
                            v44 := GenS3{F0: 20, F1: 466, F2: 0}
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                    }
                    _ = v24
                }
            default:
                v1 -= ((1.5 + v1) - v1)
            }
            if (!true) {
                var v45 uint64 = 631
                if ((3.14159 * 0.1) < v1) {
                    var v46 bool = (false || false)
                    _ = v46
                }
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch 742 {
                    case -621:
                        for fi47 := 0; fi47 < 6; fi47++ {
                            v1 = (0.0 / v1)
                            _ = len(v6)
                            for fi48 := 0; fi48 < 9; fi48++ {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = fi48
                            }
                            _ = fi47
                        }
                    case 32767:
                        v45 = 4294967295
                    }
                    switch v3 {
                    case "":
                        for fi49 := 0; fi49 < 18; fi49++ {
                            switch fi49 {
                            case -1:
                                fmt.Println(v3, fi49, v1)
                            case 9223372036854775807:
                                if true {
                                    fmt.Println(v45, v3)
                                } else {
                                    _ = len(v4)
                                }
                            }
                            var v50 int8 = 79
                            for fi51 := 0; fi51 < 17; fi51++ {
                                switch v3 {
                                case "\n\t":
                                    var v52 float32 = ((float32((((0.1 - 0.0) - -2.25) - (1e10 / (3.14159 * 0.1)))) / ((1.5 - 0.1) + float32(((0.0 * float64(-2.25)) / ((1.5 - float64(v1)) - 3.14159))))) / float32(float64(1e10)))
                                    _ = v52
                                case "a":
                                    fmt.Println(v5)
                                case "":
                                    v53 := GenS7{F0: 255, F1: 9223372036854775807}
                                    _ = v53
                                default:
                                    var v54 uint = 42
                                    _ = v54
                                }
                                _ = fi51
                            }
                            _ = fi49
                            _ = v50
                        }
                    case "hello":
                        if (-1 != -2147483648) {
                            fmt.Println(v7, v1)
                            for fi55 := 0; fi55 < 19; fi55++ {
                                v56 := GenS8{F0: (^(v5)), F1: 42, F2: 91}
                                _ = fi55
                                _ = v56
                            }
                        }
                    }
                } else {
                    fmt.Println(v45, v3)
                }
                _ = v45
            }
        } else {
            var v57 uint64 = 621
            if (v57 == v57) {
                var v58 int64 = -1
                _ = v58
            } else {
                fmt.Println(v5)
                v5 = v5
                var v59 string = v3
                _ = v59
            }
            var v60 uint16 = 2
            _ = v57
            _ = v60
        }
        _ = v7
    } else {
        var v61 int64 = 1
        _ = v61
    }
    if len(v6) > 0 {
        v6[0] = 296
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    v5 = (v5 >> ((42) & 63))
    v5 = v5
    var v62 uint = 726
    var v63 uint16 = 717
    var v64 uint32 = 265
    v65 := GenS1{F0: 0}
    for fi66 := 0; fi66 < 20; fi66++ {
        v67 := GenS5{F0: 2, F1: 0}
        for fi68 := 0; fi68 < 6; fi68++ {
            fmt.Println(v64, fi68)
            var v69 uint64 = 494
            _ = fi68
            _ = v69
        }
        _ = fi66
        _ = v67
    }
    fmt.Println(v63, v64, v62)
    if true {
        if true {
            for fi70 := 0; fi70 < 14; fi70++ {
                if true {
                    var v71 int8 = 1
                    _ = v71
                } else {
                    for fi72 := 0; fi72 < 19; fi72++ {
                        if (!true) {
                            v73 := GenS9{F0: v64, F1: 0}
                            fmt.Println(fi70, v1, fi70)
                            if ((v62 * (v62 * (v62 & v62))) <= 42) {
                                fmt.Println(fi72, fi72)
                            } else {
                                var v74 uint64 = 382
                                if ((!false) && true) {
                                    v75 := GenS10{F0: (v63 | 65535), F1: 42}
                                    _ = v75
                                } else {
                                    var v76 float32 = v1
                                    v62 = (v62 / (((v62 - 2)) | 1))
                                    v63 += (^(v63))
                                    _ = v76
                                }
                                switch (v3 + v3) {
                                case "a":
                                    v62 -= v62
                                case "hello":
                                    fmt.Println(v64, v62, v5)
                                default:
                                    v1 += (v1 / 0.1)
                                }
                                if (0.0 < float32(3.14159)) {
                                    if true {
                                        if false {
                                            var v77 int = -9223372036854775808
                                            v63 = (^(v63))
                                            _ = v77
                                        } else {
                                            v78 := GenS11{F0: 1, F1: (^(v63))}
                                            var v79 float32 = (v1 / ((1e10 / 1.0) + 0.0))
                                            if (2147483647 <= -170) {
                                                switch fi70 {
                                                case 12:
                                                    switch fi70 {
                                                    case -489:
                                                        v1 -= -2.25
                                                    case 863:
                                                        v80 := GenS12{F0: 910}
                                                        _ = v80
                                                    case 259:
                                                        var v81 int = 32768
                                                        var v82 int8 = 0
                                                        v6 = append(v6, (^(fi70)))
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                }
                                            }
                                            _ = v78
                                            _ = v79
                                        }
                                        if false {
                                            v83 := GenS13{F0: (1 >= 0), F1: ("\n\t" + "hello")}
                                            _ = v83
                                        } else {
                                            _ = len(v4)
                                            switch (v3 + "foo bar") {
                                            case "a":
                                                v84 := GenS8{F0: 1000, F1: -1, F2: 1}
                                                fmt.Println(v3)
                                                if (((v74 << (((v62 >> (((v62 >> ((v62) & 63))) & 63))) & 63)) == v74) && (-19 > 100)) {
                                                    v64 = (v64 + (v64 % (((v64 * v64)) | 1)))
                                                }
                                                _ = v84
                                            case "foo bar":
                                                v85 := GenS6{F0: v3}
                                                _ = v85
                                            case "\n\t":
                                                _ = len(v6)
                                            default:
                                                if true {
                                                    v86 := GenS2{F0: uint8(v62), F1: 127}
                                                    if (!true) {
                                                        var v87 uint = v62
                                                        v1 = v1
                                                        var v88 int16 = -741
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                    _ = v86
                                                }
                                            }
                                            var v89 int32 = -2147483648
                                            v90 := GenS4{F0: 42, F1: -9223372036854775808, F2: -1}
                                            v91 := GenS8{F0: -32768, F1: 9223372036854775807, F2: 127}
                                            var v92 uint8 = 53
                                            var v93 bool = true
                                            v1 -= (v1 / 0.1)
                                            _ = v89
                                            _ = v90
                                            _ = v91
                                            _ = v92
                                            _ = v93
                                        }
                                    } else {
                                        if true {
                                            v64 -= v64
                                            fmt.Println(fi72, v62, fi72)
                                            v94 := GenS14{F0: 44, F1: 62, F2: -83}
                                            if (fi70 < (fi70 % (((fi72 + (fi70 ^ -2147483648))) | 1))) {
                                                v95 := GenS3{F0: (v62 >> (((v62 - v62)) & 63)), F1: (v74 & (v74 / (((^(v74))) | 1))), F2: (v64 >> ((277) & 63))}
                                                if true {
                                                    var v96 int = fi72
                                                    var v97 string = ("foo bar" + "\n\t")
                                                    _ = v96
                                                    _ = v97
                                                } else {
                                                    switch (fi72 * -9223372036854775808) {
                                                    case -2147483648:
                                                        v62 = v62
                                                    }
                                                    if ((v5 < -986) || false) {
                                                        v98 := GenS14{F0: -1, F1: 42, F2: 1}
                                                        if true {
                                                            var v99 float64 = 1.5
                                                            _ = v99
                                                        } else {
                                                            v62 = uint(v63)
                                                        }
                                                        v64 = (v64 | v64)
                                                        _ = v98
                                                    }
                                                }
                                                var v100 uint16 = (v63 | 255)
                                                _ = v100
                                                _ = v95
                                            } else {
                                                switch fi72 {
                                                case 9223372036854775807:
                                                    var v101 uint16 = (v63 | 42)
                                                    _ = v101
                                                case 618:
                                                    v6 = append(v6, fi72)
                                                case -376:
                                                    var v102 int = 459
                                                    _ = v102
                                                default:
                                                    var v103 uint8 = 1
                                                    _ = v103
                                                }
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                            _ = v94
                                        } else {
                                            switch (fi72 ^ -236) {
                                            case 42:
                                                v5 += (v5 + (v5 >> (((v62 | v62)) & 63)))
                                            }
                                            if (-674 >= 2147483647) {
                                                if ((v63 * (v63 % ((uint16(v62)) | 1))) != v63) {
                                                    v63 -= v63
                                                    if (114 == int32(fi72)) {
                                                        v74 -= 750
                                                        v1 += v1
                                                    } else {
                                                        v64 += (v64 - v64)
                                                    }
                                                } else {
                                                    switch (fi72 << ((v62) & 63)) {
                                                    case -9223372036854775808:
                                                        v64 = v64
                                                    case -777:
                                                        v104 := GenS5{F0: (^(v74)), F1: 51}
                                                        _ = v104
                                                    }
                                                    if true {
                                                        v105 := GenS4{F0: -9223372036854775808, F1: -1, F2: 42}
                                                        var v106 string = v3
                                                        var v107 uint32 = (v64 % ((uint32(v63)) | 1))
                                                        _ = v105
                                                        _ = v106
                                                        _ = v107
                                                    } else {
                                                        _ = len(v6)
                                                        switch "" {
                                                        case "a":
                                                            _ = len(v2)
                                                        case "hello":
                                                            var v108 uint = (v62 - (v62 & 255))
                                                            _ = v108
                                                        case "\n\t":
                                                            switch fi70 {
                                                            case 898:
                                                                if len(v4) > 0 {
                                                                    v4[0] = 1
                                                                }
                                                            case 127:
                                                                v109 := GenS15{F0: (v74 << (((v62 ^ (v62 + (v62 << ((v62) & 63))))) & 63)), F1: (v63 << (((v62 >> (((v62 - 255)) & 63))) & 63))}
                                                                _ = v109
                                                            case 733:
                                                                var v110 string = v3
                                                                _ = v110
                                                            }
                                                        default:
                                                            var v111 int16 = (-(v5))
                                                            _ = v111
                                                        }
                                                    }
                                                }
                                            }
                                            if len(v6) > 0 {
                                                v6[0] = fi70
                                            }
                                            v112 := GenS7{F0: (v63 & v63), F1: 9223372036854775807}
                                            var v113 float32 = 0.1
                                            _ = v112
                                            _ = v113
                                        }
                                        v114 := GenS1{F0: 91}
                                        if (!(true && (!((3.14159 / ((0.1 / 1.0) / (((v1 + ((v1 / 3.14159) + v1)) - (0.0 * v1)) - (v1 * v1)))) != v1)))) {
                                            if true {
                                                var v115 int64 = int64(v5)
                                                v116 := GenS3{F0: (v62 | uint(v63)), F1: 4294967295, F2: (v64 + (v64 >> (((v62 | 97)) & 63)))}
                                                _ = v115
                                                _ = v116
                                            } else {
                                                v62 = v62
                                                if (1.5 == (((1e10 * float64(v1)) + 1.5) + (1e10 * (3.14159 + (-2.25 * 3.14159))))) {
                                                    switch (fi72 / ((0) | 1)) {
                                                    case -277:
                                                        fmt.Println(v63)
                                                    case -1:
                                                        if ((v74 * v74) > v74) {
                                                            var v117 uint = (v62 << ((uint(v64)) & 63))
                                                            _ = v117
                                                        }
                                                    }
                                                }
                                                if false {
                                                    v118 := GenS10{F0: 32767, F1: 127}
                                                    _ = v118
                                                }
                                                if true {
                                                    if (-1 >= -128) {
                                                        var v119 string = "\n\t"
                                                        _ = v119
                                                    }
                                                    v3 = "\n\t"
                                                    if false {
                                                        var v120 uint16 = v63
                                                        _ = len(v2)
                                                        v1 -= 1e10
                                                        _ = v120
                                                    } else {
                                                        switch v3 {
                                                        case "foo bar":
                                                            v121 := GenS16{F0: uint16(v62), F1: v62}
                                                            _ = v121
                                                        case "":
                                                            v62 -= v62
                                                        case "hello":
                                                            fmt.Println(v3, v1, v62)
                                                        default:
                                                            switch int(v5) {
                                                            case 2147483647:
                                                                var v122 uint8 = 128
                                                                _ = v122
                                                            case 0:
                                                                v1 = 0.1
                                                            case 32768:
                                                                v3 = ("foo bar" + v3)
                                                            default:
                                                                var v123 bool = ((uint(v74) <= v62) || (v64 != 1))
                                                                _ = v123
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    if true {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v74 = (v74 / ((42) | 1))
                                                        var v124 int = 567
                                                        _ = v124
                                                    } else {
                                                        if false {
                                                            v1 += float32(3.14159)
                                                            var v125 uint32 = 903
                                                            var v126 int32 = -1
                                                            v63 = (^(v63))
                                                            _ = v125
                                                            _ = v126
                                                        } else {
                                                            v62 = (v62 << ((0) & 63))
                                                            v127 := GenS17{F0: -2147483648}
                                                            if true {
                                                                v128 := GenS18{F0: 1, F1: -1, F2: fi70}
                                                                var v129 int = fi70
                                                                _ = v128
                                                                _ = v129
                                                            }
                                                            _ = v127
                                                        }
                                                        switch v3 {
                                                        case "\n\t":
                                                            if ((246 == 136) && false) {
                                                                if (true || true) {
                                                                    fmt.Println(v3, v63)
                                                                }
                                                                v130 := GenS11{F0: 255, F1: v63}
                                                                switch (v3 + ((((v3 + v3) + "hello") + (v3 + v3)) + "a")) {
                                                                case "foo bar":
                                                                    var v131 int16 = (^(v5))
                                                                    _ = v131
                                                                default:
                                                                    fmt.Println(v5, v3, fi72)
                                                                }
                                                                _ = v130
                                                            }
                                                        default:
                                                            v132 := GenS19{F0: uint16(v64), F1: -1, F2: -128}
                                                            _ = v132
                                                        }
                                                        fmt.Println(v1)
                                                    }
                                                }
                                            }
                                            var v133 uint16 = v63
                                            _ = v133
                                        } else {
                                            v134 := GenS9{F0: uint32(v74), F1: 628}
                                            if true {
                                                _ = len(v6)
                                                var v135 uint64 = (^(v74))
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v135
                                            }
                                            v4 = append(v4, (v64 & (v64 / (((v64 | v64)) | 1))))
                                            _ = v134
                                        }
                                        if (-9223372036854775808 < 0) {
                                            var v136 uint32 = (v64 / (((v64 * v64)) | 1))
                                            v1 = float32((0.0 / ((1e10 - float64(-2.25)) + float64((v1 * (v1 * v1))))))
                                            v3 = (v3 + v3)
                                            _ = v136
                                        }
                                        if false {
                                            v64 -= (v64 * (^(v64)))
                                        }
                                        if true {
                                            if true {
                                                var v137 int = (fi72 % ((fi70) | 1))
                                                _ = v137
                                            }
                                        } else {
                                            v5 += v5
                                            v4 = append(v4, 276)
                                        }
                                        v64 = v64
                                        switch fi70 {
                                        case -1:
                                            var v138 float32 = v1
                                            _ = v138
                                        case 907:
                                            switch (fi70 % ((fi70) | 1)) {
                                            case 4294967295:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case 259:
                                                v4 = append(v4, (v64 ^ (v64 | v64)))
                                                var v139 int32 = 0
                                                _ = v139
                                            }
                                        default:
                                            v140 := GenS20{F0: v74}
                                            _ = v140
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v114
                                    }
                                }
                                _ = v74
                            }
                            _ = v73
                        }
                        v1 += v1
                        v141 := GenS4{F0: -9223372036854775808, F1: 9223372036854775807, F2: -128}
                        _ = fi72
                        _ = v141
                    }
                }
                _ = fi70
            }
        }
        if true {
            v142 := GenS14{F0: 121, F1: -9223372036854775808, F2: 1}
            _ = v142
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v4) > 0 {
                v4[0] = (v64 / ((uint32(v63)) | 1))
            }
        }
        var v143 uint32 = (^(v64))
        _ = v143
    }
    v3 = v3
    v144 := GenS21{F0: 0}
    v145 := GenS20{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v144
    _ = v145
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
