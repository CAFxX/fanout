package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint }
type GenS2 struct { F0 int64; F1 string; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 int32; F1 uint32; F2 int64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint32; F1 string; F2 int64 }
type GenS14 struct { F0 uint32; F1 string }

func main() {
    v1 := []uint32{}
    var v2 uint8 = 116
    v3 := []string{"hello"}
    var v4 uint = 65535
    v2 += (v2 + v2)
    for fi5 := 0; fi5 < 2; fi5++ {
        _ = len(v3)
        if true {
            fmt.Println(v4, fi5)
            for fi6 := 0; fi6 < 20; fi6++ {
                if true {
                    var v7 int16 = -32768
                    v8 := GenS1{F0: -66, F1: -111, F2: (v4 % ((v4) | 1))}
                    _ = len(v3)
                    _ = v7
                    _ = v8
                }
                if true {
                    if (1 <= 42) {
                        var v9 int = fi5
                        _ = v9
                    }
                } else {
                    fmt.Println(v4)
                }
                _ = fi6
            }
            switch (fi5 * (fi5 >> ((v4) & 63))) {
            case 32767:
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 uint32 = 551
                    v2 = v2
                    v4 = v4
                    switch "" {
                    case "foo bar":
                        v12 := GenS2{F0: -369, F1: "foo bar", F2: -824}
                        _ = v12
                    case "\n\t":
                        if true {
                            switch fi5 {
                            case -9223372036854775808:
                                var v13 bool = (uint8(v11) > v2)
                                _ = v13
                            case 32768:
                                fmt.Println(v4, fi5)
                            }
                        } else {
                            var v14 uint8 = v2
                            _ = v14
                        }
                    default:
                        var v15 bool = (32767 >= 861)
                        _ = v15
                    }
                    v11 = v11
                    _ = fi10
                    _ = v11
                }
            }
        }
        _ = fi5
    }
    v16 := GenS2{F0: 600, F1: "\n\t", F2: -9223372036854775808}
    if (!false) {
        fmt.Println(v2, v2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v17 uint = 65535
        _ = v17
    } else {
        var v18 int8 = 1
        v2 = (v2 >> (((v4 << ((v4) & 63))) & 63))
        _ = v18
    }
    v19 := GenS1{F0: 127, F1: 2147483647, F2: 408}
    for fi20 := 0; fi20 < 3; fi20++ {
        _ = len(v1)
        _ = fi20
    }
    v21 := GenS1{F0: -128, F1: 2147483647, F2: v4}
    if false {
        if (false && true) {
            v4 -= 18446744073709551615
        } else {
            if (-125 > 9223372036854775807) {
                var v22 int8 = 100
                if (2 != 42) {
                    v23 := GenS3{F0: 9223372036854775807}
                    v4 -= 138
                    _ = v23
                }
                _ = v22
            } else {
                var v24 string = "hello"
                if (2147483647 > 4294967295) {
                    v24 = v24
                }
                switch 4294967295 {
                case -1:
                    for fi25 := 0; fi25 < 2; fi25++ {
                        v26 := GenS2{F0: -1, F1: (("foo bar" + v24) + ("foo bar" + (v24 + ""))), F2: -799}
                        v27 := GenS2{F0: -467, F1: v24, F2: -1}
                        var v28 int64 = 1
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                case 128:
                    v24 = v24
                case 2147483647:
                    for fi29 := 0; fi29 < 2; fi29++ {
                        var v30 int8 = 1
                        v4 += 65535
                        var v31 uint64 = 832
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                default:
                    v24 = v24
                }
                _ = v24
            }
            var v32 uint16 = 32767
            v33 := GenS3{F0: 42}
            _ = v32
            _ = v33
        }
    } else {
        var v34 uint = (v4 & v4)
        _ = v34
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        if false {
            v2 = v2
        }
        v4 = (^(v4))
        _ = fi35
    }
    fmt.Println(v2, v2)
    v36 := GenS4{F0: 4294967295, F1: 2147483647, F2: 42}
    switch 4294967295 {
    case 32768:
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
    default:
        for fi37 := 0; fi37 < 4; fi37++ {
            var v38 uint32 = 65535
            v4 = 79
            var v39 uint64 = 2
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    var v40 uint = (^(v4))
    if ((32767 == 1) && (0 != 257)) {
        v41 := GenS2{F0: -131, F1: "", F2: 4}
        for fi42 := 0; fi42 < 2; fi42++ {
            v4 -= v4
            _ = fi42
        }
        _ = v41
    } else {
        for fi43 := 0; fi43 < 20; fi43++ {
            var v44 int32 = 1
            var v45 bool = true
            var v46 uint32 = 993
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
        v47 := GenS2{F0: -1, F1: ("foo bar" + "a"), F2: -357}
        _ = v47
    }
    v40 += (v4 << ((v4) & 63))
    for fi48 := 0; fi48 < 12; fi48++ {
        switch fi48 {
        case 2147483647:
            if false {
                for fi49 := 0; fi49 < 16; fi49++ {
                    if (("a" + "hello") <= "hello") {
                        if true {
                            if false {
                                var v50 uint = 42
                                if true {
                                    v51 := GenS4{F0: fi49, F1: 65535, F2: 2}
                                    v40 = v4
                                    var v52 int16 = int16(fi48)
                                    _ = v51
                                    _ = v52
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if (int16(v4) > -1) {
                                            fmt.Println(v4, fi48)
                                            if true {
                                                if true {
                                                    if false {
                                                        switch (fi49 << ((v40) & 63)) {
                                                        case 1:
                                                            v2 = v2
                                                        }
                                                    }
                                                    var v53 int16 = -124
                                                    fmt.Println(fi49, v53, v53)
                                                    _ = v53
                                                } else {
                                                    v40 = (v4 - v40)
                                                    switch int(v2) {
                                                    case -1:
                                                        fmt.Println(v40, v4)
                                                        v54 := GenS5{F0: 1}
                                                        v50 = (v4 >> ((v40) & 63))
                                                        _ = v54
                                                    case 42:
                                                        var v55 bool = true
                                                        _ = v55
                                                    case -495:
                                                        v50 -= 578
                                                    default:
                                                        switch ("hello" + "foo bar") {
                                                        case "foo bar":
                                                            v40 = v40
                                                        default:
                                                            if false {
                                                                v56 := GenS6{F0: 1, F1: "\n\t"}
                                                                if true {
                                                                    v57 := GenS7{F0: (fi48 * fi49)}
                                                                    v58 := GenS5{F0: 42}
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    v2 += (v2 - (v2 % ((v2) | 1)))
                                                                    v50 = v40
                                                                }
                                                                v4 += 0
                                                                switch "\n\t" {
                                                                case "a":
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        fmt.Println(fi49, fi49, v2)
                                                                    default:
                                                                        if true {
                                                                            v2 += 1
                                                                        } else {
                                                                            v4 += 0
                                                                        }
                                                                    }
                                                                case "":
                                                                    var v59 int8 = -109
                                                                    _ = v59
                                                                default:
                                                                    fmt.Println(v4, fi48)
                                                                }
                                                                var v60 uint32 = 206
                                                                var v61 float32 = (0.0 + 1e10)
                                                                _ = v56
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                        }
                                                    }
                                                    if true {
                                                        var v62 uint64 = 467
                                                        _ = v62
                                                    } else {
                                                        v2 = (v2 << ((v4) & 63))
                                                        switch fi49 {
                                                        case -1:
                                                            if len(v3) > 0 {
                                                                v3[0] = ("hello" + "\n\t")
                                                            }
                                                        default:
                                                            var v63 uint32 = 784
                                                            _ = v63
                                                        }
                                                        _ = len(v3)
                                                        v40 -= (v40 << ((v40) & 63))
                                                        if len(v3) > 0 {
                                                            v3[0] = "foo bar"
                                                        }
                                                    }
                                                    v64 := GenS2{F0: 0, F1: "\n\t", F2: 505}
                                                    _ = v64
                                                }
                                                switch (fi49 / (((fi49 ^ (fi48 * -381))) | 1)) {
                                                case 0:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                default:
                                                    if (450 < -1) {
                                                        if true {
                                                            v65 := GenS6{F0: 32767, F1: "hello"}
                                                            _ = v65
                                                        } else {
                                                            v2 = (v2 * v2)
                                                        }
                                                        v2 -= uint8(v40)
                                                        if (!((-32768 == 32767) || false)) {
                                                            fmt.Println(v4, v40, fi48)
                                                            v50 = (v4 & (v40 >> ((v50) & 63)))
                                                            fmt.Println(v2, v40)
                                                        } else {
                                                            v4 -= (v4 / (((v4 - 255)) | 1))
                                                            if (!true) {
                                                                v4 = v50
                                                                v66 := GenS8{F0: 0}
                                                                _ = v66
                                                            } else {
                                                                v67 := GenS5{F0: 127}
                                                                if len(v3) > 0 {
                                                                    v3[0] = "hello"
                                                                }
                                                                v2 -= 0
                                                                _ = v67
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    case "":
                                        var v68 uint = (v40 % ((42) | 1))
                                        _ = v68
                                    }
                                }
                                v40 = 1
                                _ = v50
                            }
                            v69 := GenS9{F0: "hello", F1: 41}
                            v40 += 1
                            _ = v69
                        } else {
                            var v70 string = "a"
                            _ = v70
                        }
                        var v71 int32 = 42
                        _ = v71
                    } else {
                        switch fi49 {
                        case 127:
                            fmt.Println(v4, fi48)
                        case -771:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            v40 -= (v40 >> ((947) & 63))
                        }
                        var v72 uint16 = 950
                        v73 := GenS6{F0: 32767, F1: "\n\t"}
                        switch fi49 {
                        case 32767:
                            if (false && ((-4 != -1) || true)) {
                                var v74 int8 = 127
                                var v75 int = (fi48 & (fi49 ^ int(v74)))
                                _ = v74
                                _ = v75
                            } else {
                                fmt.Println(fi48, v4, v40)
                                var v76 bool = (-745 >= 57)
                                _ = v76
                            }
                        case 875:
                            v40 = v40
                        case -1:
                            if (18446744073709551615 == 170) {
                                if true {
                                    switch ("\n\t" + ("a" + "")) {
                                    case "hello":
                                        fmt.Println(v2, v4, fi48)
                                    }
                                    if ("a" == ("\n\t" + "")) {
                                        if false {
                                            if false {
                                                v3 = append(v3, "hello")
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v4 -= v40
                                                v77 := GenS10{F0: 185, F1: 42, F2: 810}
                                                fmt.Println(v4)
                                                v3 = append(v3, "hello")
                                                if (!(0.0 >= (1e10 - (3.14159 * ((0.0 * float32(3.14159)) / (3.14159 / 3.14159)))))) {
                                                    v4 += 1
                                                } else {
                                                    if true {
                                                        if (true || true) {
                                                            fmt.Println(fi49)
                                                            var v78 string = ("a" + "")
                                                            v79 := GenS9{F0: v78, F1: 556}
                                                            _ = v78
                                                            _ = v79
                                                        } else {
                                                            var v80 string = ("" + "\n\t")
                                                            v72 += (v72 + (v72 + v72))
                                                            _ = v80
                                                        }
                                                        if (!(v4 <= 4294967295)) {
                                                            v81 := GenS9{F0: "foo bar", F1: -590}
                                                            var v82 int32 = 42
                                                            var v83 uint8 = 0
                                                            v4 += v4
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                        } else {
                                                            v40 += (v4 & 411)
                                                        }
                                                        var v84 uint = 65535
                                                        if (0.0 == -2.25) {
                                                            if (true && ((0.1 >= 0.0) || false)) {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                v85 := GenS2{F0: 0, F1: "hello", F2: -9223372036854775808}
                                                                _ = v85
                                                            }
                                                            var v86 bool = ((-2.25 * 1e10) > 1e10)
                                                            fmt.Println(v40, v40, v84)
                                                            _ = v86
                                                        } else {
                                                            var v87 uint64 = 18446744073709551615
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            fmt.Println(v87, v40, fi49)
                                                            _ = v87
                                                        }
                                                        fmt.Println(fi49)
                                                        fmt.Println(v2, fi49, fi49)
                                                        v4 -= (v84 + (v40 >> (((v40 + 129)) & 63)))
                                                        _ = v84
                                                    } else {
                                                        v40 = (v40 & v4)
                                                        var v88 int8 = 103
                                                        if len(v1) > 0 {
                                                            v1[0] = 42
                                                        }
                                                        _ = v88
                                                    }
                                                }
                                                _ = v77
                                            } else {
                                                var v89 uint8 = (v2 / ((2) | 1))
                                                _ = v89
                                            }
                                            var v90 int8 = 100
                                            if false {
                                                var v91 uint32 = 65535
                                                fmt.Println(fi48)
                                                _ = v91
                                            }
                                            if true {
                                                if false {
                                                    var v92 int16 = -334
                                                    _ = v92
                                                }
                                            } else {
                                                if (int64(fi49) >= 0) {
                                                    if false {
                                                        v93 := GenS2{F0: 1, F1: "foo bar", F2: 1}
                                                        _ = v93
                                                    }
                                                }
                                                if (v72 >= (v72 * 420)) {
                                                    _ = len(v1)
                                                    fmt.Println(v72, v40, v90)
                                                }
                                                if false {
                                                    if false {
                                                        switch (fi49 >> (((v40 ^ v4)) & 63)) {
                                                        case 1:
                                                            var v94 uint16 = v72
                                                            _ = v94
                                                        default:
                                                            v90 = 100
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v95 := GenS1{F0: (v90 % ((v90) | 1)), F1: -1, F2: v40}
                                                        _ = v95
                                                    }
                                                    var v96 int16 = -1
                                                    fmt.Println(v96, v2, v40)
                                                    var v97 int16 = 1
                                                    v1 = append(v1, 2)
                                                    if ("foo bar" <= "") {
                                                        var v98 float32 = 1e10
                                                        _ = v98
                                                    }
                                                    v72 = v72
                                                    _ = v96
                                                    _ = v97
                                                }
                                            }
                                            _ = v90
                                        }
                                    }
                                    var v99 int = (fi49 ^ fi48)
                                    _ = v99
                                } else {
                                    if (("" != "a") && (false && ((((v40 <= v40) || true) || (!true)) || false))) {
                                        if (false && (!true)) {
                                            _ = len(v1)
                                            v4 += (^(v4))
                                            switch (fi49 << ((v40) & 63)) {
                                            case 4294967295:
                                                v100 := GenS11{F0: 2147483647}
                                                _ = v100
                                            }
                                            if true {
                                                var v101 int8 = 0
                                                fmt.Println(v72)
                                                v102 := GenS12{F0: 0}
                                                _ = v101
                                                _ = v102
                                            } else {
                                                fmt.Println(v40, v4)
                                                var v103 uint16 = 2
                                                _ = v103
                                            }
                                            fmt.Println(v4)
                                        }
                                    } else {
                                        var v104 string = "hello"
                                        v104 = v104
                                        var v105 uint8 = (^(v2))
                                        _ = v104
                                        _ = v105
                                    }
                                }
                            } else {
                                v72 = 0
                                switch "hello" {
                                case "a":
                                    if false {
                                        switch "a" {
                                        case "":
                                            if len(v1) > 0 {
                                                v1[0] = 4294967295
                                            }
                                        }
                                        switch ("\n\t" + "hello") {
                                        case "a":
                                            var v106 int64 = 9223372036854775807
                                            _ = v106
                                        case "foo bar":
                                            switch (fi48 & fi48) {
                                            case -151:
                                                switch (fi48 - fi48) {
                                                case -397:
                                                    switch fi48 {
                                                    case -626:
                                                        _ = len(v1)
                                                    case 251:
                                                        var v107 int32 = 596
                                                        _ = v107
                                                    }
                                                case 4294967295:
                                                    v72 -= v72
                                                default:
                                                    v1 = append(v1, 1)
                                                }
                                            case 32768:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case 127:
                                                v3 = append(v3, "\n\t")
                                            }
                                            if false {
                                                var v108 int32 = -2147483648
                                                if false {
                                                    _ = len(v1)
                                                    var v109 int64 = 0
                                                    _ = v109
                                                } else {
                                                    v108 -= v108
                                                    var v110 int8 = -128
                                                    if ((^(fi49)) == fi48) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    _ = v110
                                                }
                                                _ = v108
                                            } else {
                                                var v111 int = (fi48 ^ fi49)
                                                _ = v111
                                            }
                                            fmt.Println(v72, fi48)
                                        }
                                        if true {
                                            var v112 uint = (v4 % ((v4) | 1))
                                            v72 = v72
                                            v112 = v40
                                            _ = v112
                                        }
                                    }
                                case "hello":
                                    v40 -= v4
                                default:
                                    v4 = v4
                                }
                            }
                        }
                        var v113 int64 = 170
                        _ = v113
                        _ = v72
                        _ = v73
                    }
                    var v114 uint = 18446744073709551615
                    _ = fi49
                    _ = v114
                }
            }
        case -311:
            var v115 int64 = -9223372036854775808
            _ = v115
        case -1:
            if (0 < -9223372036854775808) {
                if false {
                    var v116 float64 = 0.0
                    _ = v116
                } else {
                    var v117 float32 = (0.1 / 0.1)
                    v117 += 1.5
                    _ = v117
                }
                for fi118 := 0; fi118 < 1; fi118++ {
                    v2 -= (v2 + v2)
                    switch 127 {
                    case 127:
                        var v119 int16 = -1
                        _ = v119
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi118
                }
            }
        }
        if false {
            if true {
                fmt.Println(v2, v2, fi48)
                if true {
                    for fi120 := 0; fi120 < 12; fi120++ {
                        var v121 bool = (1000 <= 0)
                        _ = fi120
                        _ = v121
                    }
                    v122 := GenS13{F0: uint32(v2), F1: (("foo bar" + "") + ""), F2: -1}
                    for fi123 := 0; fi123 < 14; fi123++ {
                        v4 = v40
                        v4 += 255
                        v3 = append(v3, ("\n\t" + ""))
                        _ = fi123
                    }
                    _ = v122
                } else {
                    var v124 uint8 = 2
                    if true {
                        for fi125 := 0; fi125 < 2; fi125++ {
                            v124 -= (v2 >> ((v40) & 63))
                            var v126 string = "\n\t"
                            _ = fi125
                            _ = v126
                        }
                        var v127 float32 = ((float32(0.0) * 1e10) + 0.1)
                        v128 := GenS7{F0: (^(fi48))}
                        _ = v127
                        _ = v128
                    } else {
                        for fi129 := 0; fi129 < 1; fi129++ {
                            var v130 bool = (false || false)
                            if v130 {
                                v130 = v130
                                v2 = (^(v124))
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                            }
                            _ = fi129
                            _ = v130
                        }
                        v1 = append(v1, 65535)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v124
                }
            } else {
                v131 := GenS14{F0: 366, F1: ""}
                v40 -= v4
                _ = v131
            }
            v132 := GenS1{F0: 0, F1: -1, F2: 2}
            var v133 uint32 = 0
            _ = v132
            _ = v133
        } else {
            var v134 int = (fi48 ^ (-(fi48)))
            _ = v134
        }
        _ = fi48
    }
    if ((v2 / (((v2 + v2)) | 1)) <= (v2 + v2)) {
        if (!true) {
            v4 = (v4 % ((v40) | 1))
        }
    }
    for fi135 := 0; fi135 < 13; fi135++ {
        _ = len(v3)
        var v136 uint64 = 1
        _ = fi135
        _ = v136
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v40
}
