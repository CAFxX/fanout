package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 int8; F2 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 int8; F2 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 int16; F1 uint }
type GenS11 struct { F0 string; F1 string; F2 uint }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int8; F1 int8 }

func main() {
    v1 := []uint32{2147483647, 2147483647, 1, 763}
    v2 := make(map[int]bool)
    var v3 uint64 = 648
    v4 := make(map[uint32]int)
    var v5 int32 = 42
    v6 := v2[-2147483648]
    var v7 int = 9223372036854775807
    for fi8 := 0; fi8 < 12; fi8++ {
        if (!v6) {
            v3 += (^(v3))
            var v9 float64 = (3.14159 - 0.1)
            var v10 int16 = 42
            _ = v10
            _ = v9
        }
        var v11 int = fi8
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi8
        _ = v11
    }
    var v12 int32 = 0
    v5 += v12
    delete(v4, 2147483647)
    for fi13 := 0; fi13 < 5; fi13++ {
        if v6 {
            var v14 uint64 = (v3 | (v3 + (v3 * (v3 ^ v3))))
            var v15 uint16 = 42
            if (0 == -32768) {
                var v16 uint64 = (v14 - (v14 ^ (v3 / ((v14) | 1))))
                fmt.Println(v12, v5)
                _ = v16
            }
            _ = v14
            _ = v15
        } else {
            if v6 {
                if ((v3 % ((v3) | 1)) == (v3 ^ (^(v3)))) {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v5 -= v5
                        if (0.1 < float32(1e10)) {
                            var v18 float32 = float32(float64(((1.5 * (0.0 - 1e10)) * (float32(1.5) * 3.14159))))
                            _ = v18
                        } else {
                            var v19 int64 = -9223372036854775808
                            v20 := GenS1{F0: "\n\t"}
                            v5 -= (-(v12))
                            _ = v19
                            _ = v20
                        }
                        _ = len(v1)
                        _ = fi17
                    }
                    v21 := GenS2{F0: 0, F1: -9223372036854775808}
                    if len(v1) > 0 {
                        v1[0] = 797
                    }
                    _ = v21
                } else {
                    if ((v7 < (fi13 >> ((846) & 63))) || false) {
                        switch ("" + "foo bar") {
                        case "a":
                            v4[0] = -805
                        case "foo bar":
                            switch (fi13 | (fi13 * (fi13 % (((v7 ^ -966)) | 1)))) {
                            case 1:
                                v12 = v5
                            default:
                                v5 += 42
                            }
                        default:
                            v22 := GenS1{F0: "a"}
                            _ = v22
                        }
                        if (42 < 1) {
                            if ((!false) || (!false)) {
                                var v23 bool = v6
                                _ = v23
                            }
                        } else {
                            v7 = 32767
                            v5 += v12
                        }
                        switch (v7 / ((fi13) | 1)) {
                        case 4294967295:
                            v3 += v3
                        }
                    } else {
                        v6 = false
                        v3 += (v3 + (v3 - v3))
                    }
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    if (!(!false)) {
                        if false {
                            var v25 int = fi13
                            if v6 {
                                v7 -= int(v3)
                                v26 := GenS3{F0: 2147483647}
                                _ = v26
                            } else {
                                fmt.Println(fi13, v3, v7)
                                v27 := GenS3{F0: 2147483647}
                                _ = v27
                            }
                            v28 := GenS4{F0: v3, F1: 0, F2: v12}
                            switch fi24 {
                            case -487:
                                if (v6 && (false && false)) {
                                    _ = len(v1)
                                    v3 += v3
                                    var v29 int = (v25 % ((v7) | 1))
                                    if v6 {
                                        if false {
                                            if v6 {
                                                switch v25 {
                                                case -1:
                                                    if (32767 >= 255) {
                                                        fmt.Println(fi13, v5)
                                                    }
                                                case 0:
                                                    var v30 uint64 = v3
                                                    _ = v30
                                                default:
                                                    var v31 uint8 = 55
                                                    _ = v31
                                                }
                                                var v32 int = v25
                                                var v33 int16 = 1000
                                                _ = v32
                                                _ = v33
                                            } else {
                                                if v6 {
                                                    v34 := GenS3{F0: 741}
                                                    if true {
                                                        if true {
                                                            switch (v29 / (((fi24 ^ (v7 % (((fi24 & (fi13 & (v25 / ((fi24) | 1))))) | 1)))) | 1)) {
                                                            case 2147483647:
                                                                fmt.Println(fi24, v12, v25)
                                                            case 787:
                                                                v35 := GenS5{F0: 325}
                                                                v36 := GenS3{F0: 770}
                                                                _ = v35
                                                                _ = v36
                                                            case -2147483648:
                                                                if (0.0 >= (0.0 - float32((3.14159 + 1.5)))) {
                                                                    var v37 string = "foo bar"
                                                                    _ = v37
                                                                } else {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 0
                                                                    }
                                                                    var v38 int64 = 1
                                                                    var v39 float32 = 1.5
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    v3 = (v3 + v3)
                                                                    _ = len(v2)
                                                                    var v40 float32 = v39
                                                                    if (!(((v3 % ((469) | 1)) <= (v3 | (v3 & v3))) && (2 < 4294967295))) {
                                                                        v2[(v25 ^ v29)] = v6
                                                                        switch 32767 {
                                                                        case 32768:
                                                                            var v41 int32 = v5
                                                                            _ = v41
                                                                        }
                                                                    } else {
                                                                        v42 := GenS3{F0: 0}
                                                                        var v43 uint32 = 296
                                                                        v44 := GenS4{F0: (v3 ^ 377), F1: 127, F2: (v5 + 2147483647)}
                                                                        _ = v42
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                    _ = v38
                                                                    _ = v39
                                                                    _ = v40
                                                                }
                                                            }
                                                            v29 += fi24
                                                            var v45 int8 = 100
                                                            _ = v45
                                                        } else {
                                                            v5 = v5
                                                            v5 -= v5
                                                            v6 = (42 != -79)
                                                            if (116 <= 42) {
                                                                v12 = (v12 % ((v12) | 1))
                                                                if v6 {
                                                                    if v6 {
                                                                        v46 := GenS1{F0: "hello"}
                                                                        v7 = 115
                                                                        switch "a" {
                                                                        case "a":
                                                                            var v47 bool = (0 > 319)
                                                                            _ = v47
                                                                        case "hello":
                                                                            v48 := GenS6{F0: 1, F1: 42, F2: "\n\t"}
                                                                            _ = v48
                                                                        case "\n\t":
                                                                            v49 := GenS6{F0: -35, F1: -59, F2: "a"}
                                                                            _ = v49
                                                                        default:
                                                                            v1 = append(v1, 72)
                                                                            var v50 float32 = 0.1
                                                                            _ = v50
                                                                        }
                                                                        _ = v46
                                                                    }
                                                                    fmt.Println(fi24, fi13)
                                                                    delete(v2, (v25 + v29))
                                                                }
                                                            }
                                                            v3 = (v3 * v3)
                                                        }
                                                    } else {
                                                        var v51 int = (-(v7))
                                                        _ = v51
                                                    }
                                                    if v6 {
                                                        var v52 int8 = 1
                                                        var v53 uint32 = 342
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    _ = v34
                                                }
                                                v6 = true
                                            }
                                            var v54 uint16 = 587
                                            v54 -= 0
                                            _ = v54
                                        } else {
                                            if v6 {
                                                _ = len(v2)
                                            } else {
                                                switch v25 {
                                                case 32768:
                                                    if (925 <= 42) {
                                                        var v55 uint = 26
                                                        _ = v55
                                                    } else {
                                                        v1 = append(v1, 4294967295)
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v56 := v4[42]
                                                            _ = v56
                                                        case "hello":
                                                            var v57 uint8 = 0
                                                            _ = v57
                                                        }
                                                        v3 = 318
                                                        if ((!(v12 != v5)) && v6) {
                                                            _ = len(v1)
                                                        } else {
                                                            var v58 bool = ("foo bar" <= "hello")
                                                            v12 -= (v12 | v12)
                                                            _ = v58
                                                        }
                                                        v25 = v29
                                                        var v59 int = fi24
                                                        _ = v59
                                                    }
                                                case 0:
                                                    v29 = -1
                                                case 127:
                                                    fmt.Println(v6)
                                                    switch v29 {
                                                    case 42:
                                                        v25 -= v7
                                                    case 2147483647:
                                                        var v60 int8 = 42
                                                        _ = v60
                                                    case 127:
                                                        v1 = append(v1, 351)
                                                    default:
                                                        v61 := GenS5{F0: 2}
                                                        _ = v61
                                                    }
                                                    if false {
                                                        fmt.Println(v6)
                                                    }
                                                }
                                            }
                                            v5 -= (v12 << ((4294967295) & 63))
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v5 -= v12
                                            var v62 uint32 = 2
                                            _ = v62
                                        }
                                    }
                                    v63 := GenS7{F0: 103}
                                    fmt.Println(fi13, v5)
                                    _ = v29
                                    _ = v63
                                } else {
                                    switch (("" + "") + "hello") {
                                    case "hello":
                                        _ = len(v1)
                                    }
                                    var v64 int8 = 1
                                    if v6 {
                                        if (868 >= 32767) {
                                            if v6 {
                                                v3 -= 1
                                                v3 = (v3 ^ 567)
                                                v3 = (v3 << ((0) & 63))
                                            } else {
                                                v2[(fi24 ^ (fi24 >> ((42) & 63)))] = true
                                                fmt.Println(v64, v3, v64)
                                                switch "hello" {
                                                case "hello":
                                                    v65 := GenS6{F0: 127, F1: v64, F2: ""}
                                                    _ = v65
                                                case "foo bar":
                                                    var v66 uint32 = 792
                                                    switch ("foo bar" + "foo bar") {
                                                    case "a":
                                                        var v67 int32 = v5
                                                        var v68 string = ""
                                                        v69 := GenS8{F0: 545}
                                                        _ = v67
                                                        _ = v68
                                                        _ = v69
                                                    case "hello":
                                                        v70 := GenS8{F0: -965}
                                                        v5 = v12
                                                        _ = v70
                                                    }
                                                    _ = v66
                                                case "":
                                                    v7 = v7
                                                default:
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v71 int = (v25 + (^(v7)))
                                                        _ = v71
                                                    case "hello":
                                                        if true {
                                                            var v72 uint64 = (v3 | 2)
                                                            v4[621] = (fi24 << ((18446744073709551615) & 63))
                                                            v2[(v25 * v7)] = (1 != 1000)
                                                            if v6 {
                                                                if v6 {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        var v73 uint32 = 2147483647
                                                                        _ = v73
                                                                    case "":
                                                                        var v74 float32 = -2.25
                                                                        _ = v74
                                                                    case "\n\t":
                                                                        if (!v6) {
                                                                            v75 := v2[fi24]
                                                                            v76 := GenS9{F0: -775, F1: 2147483648}
                                                                            _ = v75
                                                                            _ = v76
                                                                        }
                                                                    default:
                                                                        v72 = v3
                                                                        v64 += 100
                                                                    }
                                                                    var v77 float32 = 3.14159
                                                                    _ = v77
                                                                } else {
                                                                    v78 := GenS8{F0: -98}
                                                                    fmt.Println(v64, fi24)
                                                                    _ = v78
                                                                }
                                                            } else {
                                                                v6 = v6
                                                                var v79 int32 = v12
                                                                v12 -= (v12 | v12)
                                                                _ = v79
                                                            }
                                                            v3 -= v3
                                                            _ = v72
                                                        } else {
                                                            v7 = fi13
                                                            fmt.Println(v5, fi24)
                                                        }
                                                    case "":
                                                        switch (-(fi13)) {
                                                        case 42:
                                                            var v80 uint64 = 2
                                                            _ = v80
                                                        default:
                                                            v1 = append(v1, 199)
                                                        }
                                                    default:
                                                        switch "\n\t" {
                                                        case "":
                                                            v25 = 42
                                                        case "\n\t":
                                                            v6 = v6
                                                        }
                                                    }
                                                }
                                            }
                                            v5 += (v12 - v5)
                                            v81 := GenS10{F0: 42, F1: 770}
                                            if (v25 >= (v25 ^ (fi13 / ((v25) | 1)))) {
                                                v25 -= (v25 ^ v25)
                                            }
                                            _ = v81
                                        }
                                        v25 -= v7
                                        v82 := GenS11{F0: "hello", F1: "hello", F2: 896}
                                        _ = v82
                                    }
                                    var v83 uint32 = 42
                                    _ = len(v4)
                                    _ = v64
                                    _ = v83
                                }
                            }
                            _ = v25
                            _ = v28
                        } else {
                            if true {
                                switch (v7 + fi24) {
                                case 480:
                                    v3 = 4294967295
                                case 587:
                                    v5 = (v5 + v12)
                                default:
                                    v84 := GenS12{F0: -9223372036854775808, F1: -54}
                                    _ = v84
                                }
                                if v6 {
                                    var v85 uint32 = 2
                                    switch "" {
                                    case "":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "\n\t":
                                        var v86 uint64 = 595
                                        _ = v86
                                    default:
                                        if ("foo bar" == ("a" + "foo bar")) {
                                            var v87 int16 = 0
                                            fmt.Println(v85, v87, fi24)
                                            _ = v87
                                        } else {
                                            fmt.Println(v5, v12)
                                            var v88 int32 = v12
                                            _ = v88
                                        }
                                    }
                                    _ = v85
                                }
                            }
                            var v89 float32 = float32((float64(float32(1e10)) / 1.0))
                            var v90 uint32 = 4294967295
                            var v91 int = v7
                            if (3.14159 != (0.0 / 3.14159)) {
                                var v92 float64 = 1e10
                                v7 = (fi13 << ((65535) & 63))
                                _ = v92
                            }
                            v3 = 1
                            var v93 int32 = (-(v5))
                            _ = v89
                            _ = v90
                            _ = v91
                            _ = v93
                        }
                    }
                    _ = fi24
                }
            } else {
                v94 := GenS4{F0: v3, F1: -1, F2: v12}
                v95 := GenS13{F0: -115, F1: 0, F2: 1}
                _ = v94
                _ = v95
            }
            if ((v5 | 0) <= (^(v5))) {
                switch fi13 {
                case 32767:
                    var v96 int32 = 42
                    _ = v96
                case -606:
                    if v6 {
                        v7 = v7
                    }
                case 443:
                    fmt.Println(v5, v3, v3)
                }
            }
            switch -2147483648 {
            case 1:
                var v97 int32 = (v12 & (v5 - 42))
                _ = v97
            case -202:
                fmt.Println(v7, v3, v7)
            default:
                delete(v2, fi13)
            }
        }
        if (271 < 502) {
            if ((v3 <= 1) || ((v7 | (v7 - fi13)) <= 1)) {
                v4[298] = 42
                if v6 {
                    for fi98 := 0; fi98 < 4; fi98++ {
                        fmt.Println(fi98, fi98, fi13)
                        switch "" {
                        case "foo bar":
                            switch (^(fi98)) {
                            case 42:
                                v6 = (float32((-2.25 / 0.1)) < float32(0.0))
                            default:
                                if (42 <= (v5 << ((628) & 63))) {
                                    var v99 int32 = (v12 * 55)
                                    v12 -= (v99 / ((-2147483648) | 1))
                                    _ = v99
                                } else {
                                    var v100 string = ("foo bar" + ("" + ""))
                                    if v6 {
                                        delete(v2, (-(fi98)))
                                        v6 = v6
                                        fmt.Println(v100, v5)
                                    } else {
                                        v3 = (v3 | (v3 - v3))
                                    }
                                    _ = v100
                                }
                            }
                        case "\n\t":
                            v3 = (^(v3))
                        }
                        _ = fi98
                    }
                    var v101 int32 = -374
                    _ = v101
                } else {
                    v102 := GenS6{F0: -128, F1: -39, F2: ("a" + "a")}
                    fmt.Println(fi13, v12, v3)
                    _ = v102
                }
            } else {
                var v103 bool = false
                if v6 {
                    v104 := GenS10{F0: int16(v3), F1: 4294967295}
                    v105 := GenS14{F0: 1}
                    _ = v104
                    _ = v105
                } else {
                    var v106 string = "foo bar"
                    if ((2 < 219) || (!(1 <= 1000))) {
                        var v107 int32 = (v12 * v5)
                        switch (v106 + "a") {
                        case "a":
                            var v108 uint32 = 1
                            _ = v108
                        }
                        _ = v107
                    } else {
                        v106 = v106
                    }
                    v103 = v103
                    _ = v106
                }
                for fi109 := 0; fi109 < 5; fi109++ {
                    var v110 int32 = (-(v12))
                    v111 := GenS6{F0: 100, F1: 1, F2: "foo bar"}
                    _ = fi109
                    _ = v110
                    _ = v111
                }
                _ = v103
            }
            var v112 float32 = (0.0 + (1.5 * 0.1))
            switch (v7 - (fi13 << ((526) & 63))) {
            case 0:
                if (3.14159 != (3.14159 / 1e10)) {
                    for fi113 := 0; fi113 < 15; fi113++ {
                        _ = len(v2)
                        v114 := GenS15{F0: -1, F1: 42, F2: 2147483648}
                        var v115 uint = 255
                        _ = fi113
                        _ = v114
                        _ = v115
                    }
                    v12 = v5
                } else {
                    v5 += v5
                }
            case -9223372036854775808:
                v116 := GenS16{F0: (v12 <= (-(v12)))}
                _ = v116
            case 815:
                var v117 int8 = 42
                _ = v117
            }
            _ = v112
        } else {
            var v118 uint8 = 0
            v7 = fi13
            _ = v118
        }
        v2[v7] = (!v6)
        _ = fi13
    }
    for fi119 := 0; fi119 < 2; fi119++ {
        if ((0.1 - 1e10) != (1e10 * 0.0)) {
            if (!v6) {
                v5 -= v5
                v120 := GenS14{F0: 0}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v120
            } else {
                switch 908 {
                case 4294967295:
                    if (678 != 0) {
                        var v121 uint64 = v3
                        fmt.Println(v6)
                        for fi122 := 0; fi122 < 10; fi122++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi122
                        }
                        _ = v121
                    } else {
                        v12 += -1
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v1 = append(v1, 0)
                    }
                default:
                    v4[2147483647] = v7
                }
                var v123 uint = 930
                v5 += v12
                _ = v123
            }
            if v6 {
                fmt.Println(v3, fi119, v6)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v12 = 0
                delete(v2, int(v12))
                fmt.Println(v6, v7)
            }
            var v124 int = int(v3)
            _ = v124
        } else {
            switch 9223372036854775807 {
            case 366:
                v7 -= int(v12)
            case -2147483648:
                if (fi119 != (fi119 << ((892) & 63))) {
                    if (32767 >= -1) {
                        switch (fi119 & (v7 ^ 401)) {
                        case -401:
                            for fi125 := 0; fi125 < 5; fi125++ {
                                v126 := GenS13{F0: 100, F1: v3, F2: 0}
                                fmt.Println(fi119, v6, v5)
                                _ = fi125
                                _ = v126
                            }
                        case -2147483648:
                            if ("a" <= "") {
                                _ = len(v4)
                                var v127 uint64 = v3
                                _ = v127
                            }
                        case 0:
                            var v128 uint32 = 1
                            _ = v128
                        default:
                            var v129 int8 = 100
                            _ = v129
                        }
                    } else {
                        v7 += -9223372036854775808
                        v7 = 127
                        if (0 == 646) {
                            if false {
                                fmt.Println(v6)
                                for fi130 := 0; fi130 < 10; fi130++ {
                                    v131 := GenS4{F0: (v3 | v3), F1: -61, F2: (v5 << ((166) & 63))}
                                    var v132 int = fi130
                                    var v133 int8 = 0
                                    var v134 string = ""
                                    _ = fi130
                                    _ = v131
                                    _ = v132
                                    _ = v133
                                    _ = v134
                                }
                                for fi135 := 0; fi135 < 6; fi135++ {
                                    v136 := GenS10{F0: 42, F1: 4294967295}
                                    var v137 uint32 = 2147483647
                                    _ = fi135
                                    _ = v136
                                    _ = v137
                                }
                            }
                        }
                    }
                    v138 := GenS17{F0: -1, F1: -128}
                    _ = v138
                }
            case 32767:
                var v139 int64 = 917
                _ = v139
            default:
                var v140 bool = (1 >= 1)
                _ = v140
            }
            _ = len(v1)
        }
        v7 = -997
        _ = fi119
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
