package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int16; F1 int64; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 int }
type GenS11 struct { F0 uint; F1 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 string; F1 uint32; F2 bool }
type GenS15 struct { F0 uint8; F1 int }
type GenS16 struct { F0 int64; F1 uint16 }
type GenS17 struct { F0 int8; F1 int64; F2 int8 }
type GenS18 struct { F0 uint }

func main() {
    var v1 int8 = -103
    var v2 int8 = v1
    v3 := make(map[int]bool)
    var v4 int32 = -2147483648
    var v5 string = (("foo bar" + "a") + "a")
    switch 2147483647 {
    case 2147483647:
        var v6 float64 = 1e10
        _ = v6
    case 4294967295:
        var v7 float64 = -2.25
        _ = v7
    default:
        switch 32767 {
        case 128:
            v8 := []int{32768}
            _ = v8
        case 9223372036854775807:
            if (1 == 255) {
                fmt.Println(v2)
                switch ("a" + "a") {
                case "":
                    var v9 uint32 = 4294967295
                    _ = v9
                case "a":
                    var v10 int8 = 42
                    _ = v10
                case "foo bar":
                    if true {
                        v11 := GenS1{F0: 0, F1: 169}
                        var v12 int64 = -456
                        v2 -= (v2 - v1)
                        _ = v11
                        _ = v12
                    }
                }
            }
        }
    }
    fmt.Println(v4, v1)
    if (1 == -1) {
        switch 127 {
        case 508:
            switch 813 {
            case 0:
                if true {
                    fmt.Println(v2, v1)
                    var v13 uint = 988
                    _ = v13
                } else {
                    fmt.Println(v4, v1)
                    var v14 string = ""
                    v14 = ""
                    _ = v14
                }
            }
        case -2147483648:
            delete(v3, 77)
        default:
            var v15 int16 = 42
            _ = v15
        }
        v5 = v5
    }
    switch v5 {
    case "\n\t":
        v2 += (-(v1))
    default:
        if true {
            for fi16 := 0; fi16 < 10; fi16++ {
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 string = "hello"
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
            v2 = v1
            for fi19 := 0; fi19 < 19; fi19++ {
                if (((3.14159 * -2.25) <= -2.25) && (-1 > -447)) {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        v21 := GenS1{F0: 1, F1: 1}
                        v22 := GenS1{F0: -9223372036854775808, F1: 963}
                        var v23 int = (-(fi19))
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    v1 = (v2 | (^(v1)))
                }
                var v24 float32 = -2.25
                _ = fi19
                _ = v24
            }
        }
    }
    v25 := []int{0, -339, 0}
    var v26 uint64 = 4294967295
    v26 += 18446744073709551615
    if true {
        switch 32768 {
        case 1:
            switch -9223372036854775808 {
            case -2147483648:
                v27 := GenS2{F0: 0, F1: 9223372036854775807, F2: 359}
                _ = v27
            default:
                v4 = v4
            }
        case 32768:
            if (!false) {
                if ((353 <= 623) && false) {
                    v1 += (v1 / ((v1) | 1))
                    fmt.Println(v26)
                    for fi28 := 0; fi28 < 6; fi28++ {
                        if (!false) {
                            v29 := v3[(^(fi28))]
                            _ = v29
                        }
                        var v30 uint = 822
                        _ = fi28
                        _ = v30
                    }
                } else {
                    fmt.Println(v26, v26)
                }
                switch -7 {
                case -710:
                    var v31 int64 = -295
                    _ = v31
                case -327:
                    v5 = (((("a" + v5) + ((v5 + "foo bar") + v5)) + (v5 + v5)) + v5)
                case 127:
                    v32 := GenS3{F0: (v4 << ((4294967295) & 63))}
                    _ = v32
                }
                fmt.Println(v4, v4, v2)
            } else {
                switch -1 {
                case 286:
                    v26 = 18446744073709551615
                }
            }
        case -509:
            if (!false) {
                if (false || false) {
                    if (!(!true)) {
                        for fi33 := 0; fi33 < 7; fi33++ {
                            var v34 float64 = 1e10
                            _ = fi33
                            _ = v34
                        }
                        fmt.Println(v26)
                    } else {
                        for fi35 := 0; fi35 < 16; fi35++ {
                            v36 := v3[fi35]
                            _ = fi35
                            _ = v36
                        }
                        v37 := GenS1{F0: -9223372036854775808, F1: v26}
                        _ = v37
                    }
                } else {
                    v25 = append(v25, -1)
                    if (-192 >= 4294967295) {
                        for fi38 := 0; fi38 < 3; fi38++ {
                            var v39 int8 = (v2 + 120)
                            if true {
                                for fi40 := 0; fi40 < 1; fi40++ {
                                    _ = len(v25)
                                    var v41 uint = 407
                                    _ = fi40
                                    _ = v41
                                }
                                if false {
                                    fmt.Println(v5, v1, v1)
                                    v4 += v4
                                    var v42 uint32 = 65535
                                    _ = v42
                                }
                            }
                            v5 = (v5 + ("" + "a"))
                            _ = fi38
                            _ = v39
                        }
                        var v43 bool = true
                        var v44 uint64 = (v26 | (^(v26)))
                        _ = v43
                        _ = v44
                    }
                }
                if ((uint32(v26) != 2) && (true || true)) {
                    for fi45 := 0; fi45 < 19; fi45++ {
                        if (128 == 42) {
                            fmt.Println(v1, v4)
                            v4 -= (v4 << ((1) & 63))
                        } else {
                            var v46 int = (fi45 >> ((703) & 63))
                            _ = v46
                        }
                        _ = fi45
                    }
                    var v47 uint64 = (v26 % ((4294967295) | 1))
                    _ = v47
                }
            } else {
                v48 := v3[-1]
                for fi49 := 0; fi49 < 19; fi49++ {
                    if true {
                        v4 -= (v4 << ((18446744073709551615) & 63))
                    } else {
                        v50 := GenS4{F0: (v26 / (((v26 - v26)) | 1)), F1: 42, F2: 0}
                        _ = v50
                    }
                    v4 = (v4 & -1)
                    _ = fi49
                }
                fmt.Println(v2, v48, v2)
                _ = v48
            }
        default:
            fmt.Println(v2, v1, v26)
        }
        for fi51 := 0; fi51 < 14; fi51++ {
            v5 = "foo bar"
            fmt.Println(v1, fi51, v5)
            for fi52 := 0; fi52 < 18; fi52++ {
                if false {
                    var v53 float32 = (1.5 / -2.25)
                    var v54 string = ("" + "hello")
                    var v55 float32 = v53
                    _ = v53
                    _ = v54
                    _ = v55
                } else {
                    v56 := GenS5{F0: 2, F1: 59}
                    switch 4294967295 {
                    case 127:
                        if true {
                            if (!false) {
                                v57 := GenS5{F0: 56, F1: -119}
                                switch fi52 {
                                case -9223372036854775808:
                                    if (0 > 0) {
                                        v3[-2147483648] = (v1 == v1)
                                        fmt.Println(fi52, fi52)
                                        if (false || false) {
                                            if ((134 >= 65535) && ((false && false) || (!true))) {
                                                switch (("hello" + v5) + v5) {
                                                case "hello":
                                                    v58 := GenS1{F0: 42, F1: 1}
                                                    fmt.Println(v1, fi51, fi52)
                                                    _ = v58
                                                case "":
                                                    v1 += int8(v4)
                                                case "\n\t":
                                                    var v59 int32 = (v4 | v4)
                                                    _ = v59
                                                default:
                                                    if (105 > 255) {
                                                        switch "" {
                                                        case "\n\t":
                                                            fmt.Println(v1, fi51)
                                                        case "a":
                                                            v26 = (v26 & v26)
                                                        case "hello":
                                                            var v60 bool = false
                                                            _ = v60
                                                        }
                                                        var v61 float32 = 1e10
                                                        _ = v61
                                                    }
                                                }
                                                v62 := GenS2{F0: 0, F1: 0, F2: 1}
                                                if false {
                                                    v63 := GenS6{F0: 2, F1: v1, F2: uint32(v26)}
                                                    _ = v63
                                                } else {
                                                    var v64 float64 = 0.1
                                                    _ = v64
                                                }
                                                _ = v62
                                            }
                                            v4 += v4
                                            v4 = v4
                                        } else {
                                            var v65 string = ((v5 + "a") + "foo bar")
                                            _ = v65
                                        }
                                    } else {
                                        _ = len(v25)
                                    }
                                case 9223372036854775807:
                                    var v66 float64 = (float64((1e10 * 3.14159)) + 0.0)
                                    _ = v66
                                default:
                                    switch fi52 {
                                    case -864:
                                        var v67 int32 = v4
                                        _ = v67
                                    case 2147483647:
                                        v26 = 42
                                    }
                                }
                                _ = v57
                            }
                            v68 := GenS6{F0: 2, F1: (v1 << ((42) & 63)), F2: 2147483647}
                            switch fi52 {
                            case 292:
                                fmt.Println(v2, v2, v2)
                            case -9223372036854775808:
                                v26 = v26
                                var v69 int = 42
                                v4 -= (v4 | v4)
                                _ = v69
                            case 128:
                                fmt.Println(v4, v1)
                            default:
                                v5 = (v5 + v5)
                            }
                            var v70 uint = 42
                            switch (fi51 ^ (^(fi52))) {
                            case -145:
                                if (-32768 > -1) {
                                    if (!(false || false)) {
                                        fmt.Println(v4)
                                    } else {
                                        v26 = 18446744073709551615
                                    }
                                    if (!(true && (32767 < 65535))) {
                                        v4 -= v4
                                        var v71 int = 32768
                                        if len(v25) >= 2 {
                                            v25 = v25[1:len(v25)]
                                        }
                                        v26 = v26
                                        if true {
                                            if true {
                                                v4 = (v4 & (^(v4)))
                                            }
                                            if (179 > 2147483648) {
                                                switch v5 {
                                                case "hello":
                                                    switch "" {
                                                    case "foo bar":
                                                        var v72 float32 = 1.5
                                                        _ = v72
                                                    case "a":
                                                        v73 := GenS7{F0: 2147483647, F1: (v1 % (((v1 ^ v1)) | 1))}
                                                        _ = v73
                                                    case "":
                                                        v71 -= (fi51 + (fi52 / ((fi51) | 1)))
                                                    default:
                                                        v74 := GenS8{F0: (v4 * (v4 ^ (v4 / ((v4) | 1))))}
                                                        fmt.Println(v70, fi51, v71)
                                                        _ = v74
                                                    }
                                                case "":
                                                    if len(v25) >= 2 {
                                                        v25 = v25[1:len(v25)]
                                                    }
                                                default:
                                                    v1 -= v1
                                                }
                                                fmt.Println(v26, v26)
                                                v1 = 127
                                                v26 = (v26 % ((v26) | 1))
                                            }
                                        } else {
                                            v70 += 42
                                            v26 += v26
                                            _ = len(v25)
                                            switch v71 {
                                            case -23:
                                                if (!true) {
                                                    var v75 int16 = 42
                                                    if (88 < uint8(v26)) {
                                                        switch fi51 {
                                                        case 0:
                                                            v71 = -946
                                                        case -735:
                                                            v76 := GenS2{F0: v75, F1: 1, F2: 1}
                                                            _ = v76
                                                        default:
                                                            fmt.Println(v75, v26)
                                                            fmt.Println(v26)
                                                            fmt.Println(v4, v26, v70)
                                                            var v77 int16 = (v75 ^ v75)
                                                            _ = v77
                                                        }
                                                    }
                                                    switch (v5 + "") {
                                                    case "foo bar":
                                                        v78 := GenS4{F0: 1, F1: 354, F2: (v4 >> ((4294967295) & 63))}
                                                        _ = v78
                                                    default:
                                                        fmt.Println(v75)
                                                    }
                                                    _ = v75
                                                } else {
                                                    if (!true) {
                                                        var v79 string = ""
                                                        v80 := GenS9{F0: v70}
                                                        v2 -= (v2 % (((v1 ^ int8(v70))) | 1))
                                                        _ = v79
                                                        _ = v80
                                                    } else {
                                                        v81 := GenS10{F0: 0, F1: (fi52 * (fi52 | -1))}
                                                        _ = v81
                                                    }
                                                    if true {
                                                        if (v70 > (v70 + 4294967295)) {
                                                            if len(v25) >= 2 {
                                                                v25 = v25[1:len(v25)]
                                                            }
                                                        }
                                                    }
                                                    var v82 int8 = 42
                                                    _ = v82
                                                }
                                            case 128:
                                                v83 := GenS9{F0: v70}
                                                _ = v83
                                            case 4294967295:
                                                if false {
                                                    v3[fi51] = false
                                                    if true {
                                                        fmt.Println(v1)
                                                        if (0 == uint8(v70)) {
                                                            var v84 uint32 = 2
                                                            var v85 int32 = 876
                                                            if true {
                                                                var v86 uint16 = 32767
                                                                _ = v86
                                                            }
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                    }
                                                    switch (v5 + "foo bar") {
                                                    case "\n\t":
                                                        var v87 bool = ((v2 << ((v70) & 63)) < v2)
                                                        _ = v87
                                                    case "a":
                                                        switch v5 {
                                                        case "\n\t":
                                                            var v88 int64 = -1
                                                            _ = v88
                                                        case "foo bar":
                                                            fmt.Println(v71)
                                                        case "hello":
                                                            v89 := GenS9{F0: 65535}
                                                            _ = v89
                                                        }
                                                    case "foo bar":
                                                        v4 = v4
                                                        v1 = -128
                                                        if (false || (true || false)) {
                                                            switch (("\n\t" + ("a" + v5)) + v5) {
                                                            case "foo bar":
                                                                var v90 string = ((v5 + v5) + ((v5 + v5) + ""))
                                                                _ = v90
                                                            default:
                                                                var v91 uint64 = (v26 | (v26 - 66))
                                                                v92 := GenS11{F0: (v70 % ((v70) | 1)), F1: v5}
                                                                if len(v25) >= 2 {
                                                                    v25 = v25[1:len(v25)]
                                                                }
                                                                _ = v91
                                                                _ = v92
                                                            }
                                                        } else {
                                                            switch fi51 {
                                                            case 1:
                                                                if false {
                                                                    var v93 uint32 = 2147483647
                                                                    if (1.5 <= -2.25) {
                                                                        switch ((v5 + (v5 + v5)) + v5) {
                                                                        case "\n\t":
                                                                            fmt.Println(v93)
                                                                        case "":
                                                                            var v94 bool = true
                                                                            _ = v94
                                                                        case "foo bar":
                                                                            delete(v3, v71)
                                                                        default:
                                                                            var v95 uint64 = v26
                                                                            _ = v95
                                                                        }
                                                                        var v96 string = ("" + v5)
                                                                        v71 -= (fi51 / (((v71 ^ fi51)) | 1))
                                                                        if false {
                                                                            if ((v4 | v4) == -670) {
                                                                                if len(v25) > 0 {
                                                                                    v25[0] = fi51
                                                                                }
                                                                                if (1 >= 42) {
                                                                                    v97 := GenS12{F0: v70}
                                                                                    _ = v97
                                                                                } else {
                                                                                    v2 = v1
                                                                                    v98 := GenS13{F0: 808}
                                                                                    v25 = append(v25, fi52)
                                                                                    _ = v98
                                                                                }
                                                                                if len(v25) >= 2 {
                                                                                    v25 = v25[1:len(v25)]
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if (false && false) {
                                                                                if len(v25) >= 2 {
                                                                                    v25 = v25[1:len(v25)]
                                                                                }
                                                                                var v99 uint32 = v93
                                                                                _ = v99
                                                                            } else {
                                                                                v4 -= v4
                                                                            }
                                                                        }
                                                                        fmt.Println(v26, v2, v70)
                                                                        v100 := v3[(fi52 | int(v26))]
                                                                        if v100 {
                                                                            if len(v25) >= 2 {
                                                                                v25 = v25[1:len(v25)]
                                                                            }
                                                                        }
                                                                        v101 := GenS3{F0: -2147483648}
                                                                        fmt.Println(v5)
                                                                        v102 := GenS8{F0: (v4 >> ((v70) & 63))}
                                                                        _ = v100
                                                                        _ = v101
                                                                        _ = v102
                                                                        _ = v96
                                                                    } else {
                                                                        var v103 bool = (false || (uint(v93) >= (v70 - (v70 / ((v70) | 1)))))
                                                                        v104 := GenS14{F0: ((((v5 + v5) + ("a" + ("foo bar" + v5))) + "a") + ("\n\t" + "\n\t")), F1: v93, F2: v103}
                                                                        _ = v103
                                                                        _ = v104
                                                                    }
                                                                    _ = v93
                                                                } else {
                                                                    v26 += v26
                                                                }
                                                            default:
                                                                if len(v25) > 0 {
                                                                    v25[0] = 32767
                                                                }
                                                                if ((127 >= 1) || (1.5 < -2.25)) {
                                                                    var v105 int8 = v1
                                                                    fmt.Println(v105)
                                                                    _ = v105
                                                                }
                                                            }
                                                        }
                                                        if false {
                                                            fmt.Println(fi51, v5, v71)
                                                            var v106 int8 = v2
                                                            var v107 int32 = (v4 >> (((v70 * v70)) & 63))
                                                            var v108 string = (v5 + ("hello" + v5))
                                                            v71 += v71
                                                            _ = v106
                                                            _ = v107
                                                            _ = v108
                                                        }
                                                        v5 = v5
                                                        var v109 uint = v70
                                                        v2 = v1
                                                        v110 := GenS9{F0: (v109 % ((298) | 1))}
                                                        var v111 uint64 = (^(v26))
                                                        _ = v109
                                                        _ = v110
                                                        _ = v111
                                                    default:
                                                        v112 := GenS13{F0: 0}
                                                        _ = v112
                                                    }
                                                } else {
                                                    if true {
                                                        switch -9223372036854775808 {
                                                        case 0:
                                                            switch (v71 * fi51) {
                                                            case -820:
                                                                v26 += v26
                                                            case -56:
                                                                v2 += v1
                                                            default:
                                                                switch (fi52 % (((v71 | v71)) | 1)) {
                                                                case 42:
                                                                    v113 := GenS15{F0: 2, F1: (v71 ^ 42)}
                                                                    _ = v113
                                                                default:
                                                                    _ = len(v25)
                                                                }
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v71, v26)
                                                    switch v5 {
                                                    case "hello":
                                                        if true {
                                                            switch 127 {
                                                            case 9223372036854775807:
                                                                v114 := GenS8{F0: (-(v4))}
                                                                _ = v114
                                                            }
                                                            v1 = (v1 << ((v70) & 63))
                                                        }
                                                    case "foo bar":
                                                        var v115 uint32 = 65535
                                                        _ = v115
                                                    case "\n\t":
                                                        v116 := GenS8{F0: (-(v4))}
                                                        _ = v116
                                                    }
                                                }
                                            default:
                                                v70 = (^(v70))
                                            }
                                        }
                                        _ = v71
                                    } else {
                                        if false {
                                            var v117 float64 = 0.1
                                            _ = v117
                                        }
                                    }
                                } else {
                                    switch "a" {
                                    case "":
                                        v5 = ("\n\t" + v5)
                                    }
                                    v4 = 2147483647
                                }
                            case -2147483648:
                                switch (-(fi52)) {
                                case -9223372036854775808:
                                    v2 = 1
                                case 617:
                                    if (!true) {
                                        var v118 uint8 = 85
                                        var v119 float64 = (1.5 * -2.25)
                                        _ = v118
                                        _ = v119
                                    } else {
                                        v3[(fi51 << ((v70) & 63))] = false
                                        fmt.Println(v4)
                                    }
                                    fmt.Println(fi51)
                                }
                            default:
                                if true {
                                    if false {
                                        var v120 int64 = -9223372036854775808
                                        v26 += v26
                                        var v121 uint32 = 36
                                        _ = v120
                                        _ = v121
                                    }
                                    v122 := v3[int(v1)]
                                    var v123 string = v5
                                    _ = v122
                                    _ = v123
                                } else {
                                    var v124 bool = (true && (v4 != v4))
                                    v70 = 0
                                    _ = v124
                                }
                            }
                            var v125 int64 = 0
                            v26 = v26
                            fmt.Println(fi51, v2)
                            var v126 uint = (^(v70))
                            _ = v125
                            _ = v126
                            _ = v68
                            _ = v70
                        } else {
                            var v127 uint32 = 4294967295
                            var v128 uint = 42
                            v129 := GenS16{F0: -358, F1: 32767}
                            _ = v127
                            _ = v128
                            _ = v129
                        }
                    case 0:
                        v130 := GenS17{F0: v1, F1: 0, F2: v1}
                        _ = v130
                    case 2147483647:
                        if false {
                            fmt.Println(v5, fi51, fi51)
                            var v131 int32 = v4
                            v26 = v26
                            _ = v131
                        } else {
                            v2 = -1
                        }
                    }
                    _ = v56
                }
                switch ("foo bar" + (v5 + v5)) {
                case "foo bar":
                    var v132 float64 = (1e10 * (float64(0.0) * 0.0))
                    _ = v132
                }
                if (v5 >= "") {
                    v5 = ("foo bar" + ("hello" + v5))
                } else {
                    v133 := GenS8{F0: v4}
                    if false {
                        fmt.Println(v5)
                        if (fi51 == 0) {
                            v134 := GenS13{F0: -32768}
                            _ = v134
                        } else {
                            if true {
                                if true {
                                    if ((true && false) || false) {
                                        switch fi52 {
                                        case -95:
                                            var v135 int16 = 0
                                            _ = v135
                                        default:
                                            if len(v25) > 0 {
                                                v25[0] = (fi51 * (fi51 ^ fi52))
                                            }
                                        }
                                        var v136 uint8 = 0
                                        _ = v136
                                    }
                                    fmt.Println(v1, v1)
                                    delete(v3, (fi51 * 4294967295))
                                }
                            }
                            _ = len(v25)
                            switch (-(fi52)) {
                            case 32768:
                                if true {
                                    if (!((v26 * v26) <= 0)) {
                                        fmt.Println(v4)
                                    }
                                    v26 = 42
                                    var v137 float32 = -2.25
                                    v25 = append(v25, 708)
                                    v5 = v5
                                    var v138 int64 = 1
                                    _ = v137
                                    _ = v138
                                } else {
                                    v26 -= v26
                                }
                            case -9223372036854775808:
                                v3[fi52] = (255 <= 950)
                            case 9223372036854775807:
                                var v139 float64 = 0.0
                                switch fi51 {
                                case 32768:
                                    v26 = 1
                                case 234:
                                    var v140 uint64 = (^(v26))
                                    _ = v140
                                case -9223372036854775808:
                                    v26 = v26
                                default:
                                    v141 := GenS18{F0: 1}
                                    _ = v141
                                }
                                var v142 int64 = 9223372036854775807
                                _ = v139
                                _ = v142
                            default:
                                v2 -= v1
                            }
                            var v143 uint32 = 0
                            v5 = "a"
                            _ = len(v25)
                            if (!false) {
                                var v144 int64 = 42
                                if (126 == 106) {
                                    var v145 uint32 = v143
                                    fmt.Println(v145)
                                    _ = v145
                                }
                                _ = v144
                            }
                            _ = v143
                        }
                        switch "\n\t" {
                        case "a":
                            var v146 uint32 = 186
                            if true {
                                v5 = (("hello" + (v5 + v5)) + v5)
                                var v147 int32 = (v4 / ((-555) | 1))
                                _ = v147
                            }
                            _ = v146
                        }
                        v148 := GenS13{F0: 424}
                        switch 2147483647 {
                        case 32767:
                            var v149 string = ("\n\t" + v5)
                            _ = v149
                        default:
                            switch fi52 {
                            case 128:
                                var v150 int8 = v2
                                _ = v150
                            case -143:
                                _ = len(v25)
                            }
                        }
                        _ = v148
                    }
                    _ = v133
                }
                fmt.Println(v1, v4, v4)
                if true {
                    var v151 bool = false
                    var v152 uint = 0
                    fmt.Println(v152, v1, v5)
                    v153 := v3[(fi52 / ((fi51) | 1))]
                    _ = v151
                    _ = v152
                    _ = v153
                } else {
                    if (4294967295 < 786) {
                        _ = len(v25)
                    }
                    fmt.Println(v2, v26)
                }
                _ = fi52
            }
            _ = fi51
        }
    }
    v25 = append(v25, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
