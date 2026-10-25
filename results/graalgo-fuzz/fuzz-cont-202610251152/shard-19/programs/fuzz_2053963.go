package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 string; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 uint8; F2 bool }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 int8; F1 int; F2 uint16 }
type GenS11 struct { F0 uint8; F1 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint; F1 int; F2 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-828}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    switch -6 {
    case 127:
        var v5 bool = (255 > 42)
        _ = v5
    case -49:
        var v6 float32 = 1.5
        _ = v6
    case 128:
        if (42 >= 63) {
            if false {
                v7 := GenS1{F0: -1}
                if false {
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v8 int16 = -32768
                        if (42 > 1) {
                            switch ((("\n\t" + (("a" + "") + "a")) + ("hello" + ("a" + ("foo bar" + (("\n\t" + "hello") + "hello"))))) + "hello") {
                            case "\n\t":
                                v9 := GenS1{F0: (v8 + v8)}
                                _ = v9
                            case "a":
                                v8 = v8
                            default:
                                if ("" < "foo bar") {
                                    var v10 int64 = -1
                                    v10 += (v10 << ((1) & 63))
                                    v10 = -670
                                    _ = v10
                                }
                            }
                            for fi11 := 0; fi11 < 4; fi11++ {
                                if (-9223372036854775808 > 1) {
                                    v12 := GenS2{F0: -32768, F1: "hello", F2: 42}
                                    _ = v12
                                } else {
                                    v8 += v8
                                }
                                switch "a" {
                                case "a":
                                    if (false && false) {
                                        _ = len(v2)
                                    }
                                case "foo bar":
                                    fmt.Println(v8, v8)
                                }
                                v8 -= v8
                                _ = fi11
                            }
                        } else {
                            if (true && true) {
                                v8 += (^(v8))
                                for fi13 := 0; fi13 < 15; fi13++ {
                                    for fi14 := 0; fi14 < 18; fi14++ {
                                        switch "\n\t" {
                                        case "a":
                                            var v15 uint32 = 243
                                            _ = v15
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = fi14
                                            }
                                        }
                                        fmt.Println(fi13, v8, fi13)
                                        if (32767 <= 783) {
                                            v8 = 1000
                                            var v16 uint64 = 321
                                            _ = v16
                                        }
                                        _ = fi14
                                    }
                                    v8 += 200
                                    _ = fi13
                                }
                            }
                            var v17 int8 = 42
                            var v18 int = -151
                            _ = v17
                            _ = v18
                        }
                        _ = v8
                    }
                    _ = len(v4)
                } else {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        var v20 int = 9223372036854775807
                        _ = fi19
                        _ = v20
                    }
                    if ((!false) || (-1 == 2147483647)) {
                        if true {
                            v21 := GenS3{F0: 0}
                            v22 := GenS3{F0: -590}
                            _ = v21
                            _ = v22
                        }
                        var v23 uint64 = 268
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v23
                    }
                    v24 := GenS2{F0: 32767, F1: "hello", F2: 42}
                    _ = v24
                }
                _ = v7
            }
            for fi25 := 0; fi25 < 7; fi25++ {
                v26 := GenS3{F0: -2147483648}
                _ = fi25
                _ = v26
            }
            var v27 int16 = 1000
            _ = v27
        } else {
            var v28 int64 = -1
            if (461 >= -32768) {
                var v29 uint32 = 1
                if false {
                    var v30 int16 = 601
                    v31 := GenS1{F0: 0}
                    v28 = 42
                    _ = v30
                    _ = v31
                } else {
                    v29 = 1
                }
                _ = v29
            }
            _ = v28
        }
    default:
        v32 := GenS2{F0: 42, F1: ((("hello" + "") + ("" + "")) + "a"), F2: 4294967295}
        _ = v32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        v34 := GenS3{F0: 42}
        _ = fi33
        _ = v34
    }
    switch 340 {
    case 42:
        var v35 bool = false
        _ = v35
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        switch ("" + "\n\t") {
        case "\n\t":
            for fi37 := 0; fi37 < 5; fi37++ {
                var v38 int16 = -32768
                v38 = v38
                v38 = v38
                _ = len(v4)
                _ = fi37
                _ = v38
            }
        case "a":
            var v39 uint16 = 453
            _ = v39
        }
        var v40 int32 = 342
        v40 = (v40 & -1)
        _ = fi36
        _ = v40
    }
    switch "\n\t" {
    case "\n\t":
        v41 := GenS2{F0: 32767, F1: "", F2: 271}
        _ = v41
    default:
        for fi42 := 0; fi42 < 19; fi42++ {
            fmt.Println(fi42)
            for fi43 := 0; fi43 < 14; fi43++ {
                if true {
                    var v44 int16 = 0
                    _ = v44
                }
                var v45 uint64 = 2
                fmt.Println(fi42, fi43)
                _ = fi43
                _ = v45
            }
            var v46 int16 = 635
            _ = fi42
            _ = v46
        }
    }
    v47 := GenS4{F0: "a"}
    var v48 string = "\n\t"
    v49 := GenS4{F0: ((v48 + "a") + ((v48 + v48) + (v48 + ("\n\t" + v48))))}
    var v50 int8 = 0
    if (-9223372036854775808 < -667) {
        var v51 int64 = -9223372036854775808
        var v52 uint8 = 2
        _ = v51
        _ = v52
    }
    fmt.Println(v50)
    _ = len(v2)
    v50 = (v50 * (v50 << ((65535) & 63)))
    v53 := GenS5{F0: 0, F1: (v48 + ((v48 + v48) + "foo bar"))}
    if (v48 == "") {
        if (v50 >= (v50 / ((v50) | 1))) {
            switch v48 {
            case "\n\t":
                if (v48 != (v48 + (v48 + "a"))) {
                    if false {
                        v3[-762] = 4294967295
                        switch (v48 + v48) {
                        case "hello":
                            for fi54 := 0; fi54 < 2; fi54++ {
                                v50 += (v50 + v50)
                                _ = fi54
                            }
                        case "\n\t":
                            for fi55 := 0; fi55 < 5; fi55++ {
                                for fi56 := 0; fi56 < 11; fi56++ {
                                    var v57 int16 = 288
                                    if ((0 == 255) && (42 != 18446744073709551615)) {
                                        if (((v48 + v48) == v48) || true) {
                                            switch v48 {
                                            case "\n\t":
                                                var v58 int8 = (v50 + (v50 + v50))
                                                _ = v58
                                            default:
                                                fmt.Println(v50)
                                            }
                                        }
                                        v57 = (v57 & v57)
                                        switch fi55 {
                                        case 32768:
                                            if false {
                                                v1[((v48 + (v48 + "\n\t")) + "a")] = 887
                                                var v59 int32 = -545
                                                var v60 int64 = -9223372036854775808
                                                v4[209] = "hello"
                                                v61 := GenS6{F0: 53}
                                                _ = v59
                                                _ = v60
                                                _ = v61
                                            } else {
                                                v50 -= -128
                                                switch (v48 + "") {
                                                case "foo bar":
                                                    v62 := GenS2{F0: v57, F1: (v48 + "a"), F2: 42}
                                                    _ = v62
                                                default:
                                                    switch fi55 {
                                                    case 32767:
                                                        var v63 string = ("a" + v48)
                                                        _ = v63
                                                    case -2147483648:
                                                        fmt.Println(v48)
                                                    case 9223372036854775807:
                                                        v48 = v48
                                                    default:
                                                        v57 = (v57 - 1000)
                                                    }
                                                }
                                                if false {
                                                    v57 = v57
                                                    fmt.Println(fi56, fi56, v57)
                                                    v64 := v3[(fi55 & fi55)]
                                                    _ = v64
                                                } else {
                                                    v65 := GenS7{F0: "\n\t"}
                                                    _ = v65
                                                }
                                                v50 = (^(v50))
                                                var v66 int8 = (v50 + 1)
                                                v67 := GenS7{F0: v48}
                                                _ = v66
                                                _ = v67
                                            }
                                        default:
                                            v57 -= (v57 & v57)
                                        }
                                    } else {
                                        if false {
                                            v48 = "a"
                                            if false {
                                                v48 = (((v48 + v48) + (((v48 + v48) + (v48 + v48)) + ("foo bar" + v48))) + v48)
                                                v68 := GenS3{F0: 2147483647}
                                                _ = v68
                                            }
                                        } else {
                                            var v69 int16 = (-(v57))
                                            _ = v69
                                        }
                                        if (false && ((false || true) && false)) {
                                            v3[42] = 4294967295
                                            var v70 uint16 = 540
                                            _ = v70
                                        }
                                    }
                                    _ = fi56
                                    _ = v57
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if false {
                                    switch fi55 {
                                    case 938:
                                        var v71 uint64 = 1
                                        _ = v71
                                    default:
                                        v48 = "\n\t"
                                    }
                                    if false {
                                        if false {
                                            v72 := v1[((v48 + "\n\t") + ((v48 + v48) + (v48 + "\n\t")))]
                                            if false {
                                                var v73 uint = 403
                                                delete(v3, int(v73))
                                                for fi74 := 0; fi74 < 6; fi74++ {
                                                    v75 := GenS3{F0: 42}
                                                    v73 -= 420
                                                    _ = fi74
                                                    _ = v75
                                                }
                                                _ = v73
                                            }
                                            _ = v72
                                        }
                                    } else {
                                        if (v48 != (v48 + v48)) {
                                            for fi76 := 0; fi76 < 17; fi76++ {
                                                switch fi76 {
                                                case -947:
                                                    _ = len(v1)
                                                case 4294967295:
                                                    if false {
                                                        if false {
                                                            var v77 float32 = (1.5 * 0.0)
                                                            _ = v77
                                                        }
                                                        v50 = v50
                                                        if false {
                                                            v48 = v48
                                                            if (true || false) {
                                                                fmt.Println(fi76, v48, v50)
                                                            }
                                                            v50 -= v50
                                                        } else {
                                                            if false {
                                                                v50 -= 68
                                                            } else {
                                                                v78 := GenS8{F0: 564, F1: 0, F2: false}
                                                                switch (fi55 ^ fi55) {
                                                                case -409:
                                                                    var v79 float64 = (1e10 - -2.25)
                                                                    var v80 int16 = 0
                                                                    if (1 <= 18446744073709551615) {
                                                                        v81 := GenS9{F0: 128, F1: fi76}
                                                                        _ = v81
                                                                    }
                                                                    _ = v79
                                                                    _ = v80
                                                                case 1:
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                case 127:
                                                                    var v82 int32 = 0
                                                                    _ = v82
                                                                }
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v83 int16 = -93
                                                                    _ = v83
                                                                default:
                                                                    var v84 uint16 = 65535
                                                                    _ = v84
                                                                }
                                                                _ = v78
                                                            }
                                                            if (true || false) {
                                                                v50 += v50
                                                                if (("hello" + v48) <= v48) {
                                                                    if (4294967295 >= 739) {
                                                                        if false {
                                                                            fmt.Println(fi76, fi76)
                                                                            switch v48 {
                                                                            case "\n\t":
                                                                                var v85 int64 = -1
                                                                                _ = v85
                                                                            default:
                                                                                if (0 < 1) {
                                                                                    switch v48 {
                                                                                    case "a":
                                                                                        fmt.Println(fi76, v48, fi55)
                                                                                    case "foo bar":
                                                                                        if false {
                                                                                            v86 := GenS10{F0: (v50 >> ((182) & 63)), F1: fi55, F2: 2}
                                                                                            _ = v86
                                                                                        } else {
                                                                                            if true {
                                                                                                fmt.Println(v50, fi55)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if true {
                                                                                        switch "" {
                                                                                        case "":
                                                                                            var v87 int16 = 306
                                                                                            _ = v87
                                                                                        case "\n\t":
                                                                                            var v88 int16 = 1000
                                                                                            _ = v88
                                                                                        default:
                                                                                            fmt.Println(v48, fi55)
                                                                                        }
                                                                                    }
                                                                                    v89 := GenS11{F0: 1, F1: 102}
                                                                                    var v90 int = (fi76 % (((fi55 ^ -9223372036854775808)) | 1))
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if (0 != -8) {
                                                                            if (((1e10 + 1e10) <= -2.25) && false) {
                                                                                v50 = v50
                                                                            }
                                                                        }
                                                                        v50 = v50
                                                                        v91 := GenS12{F0: (v50 / (((v50 % ((v50) | 1))) | 1))}
                                                                        var v92 int = (^(fi55))
                                                                        _ = v91
                                                                        _ = v92
                                                                    }
                                                                }
                                                            }
                                                            if (!false) {
                                                                var v93 int = (fi76 & (fi55 - fi76))
                                                                if len(v2) > 0 {
                                                                    v2[0] = 182
                                                                }
                                                                _ = v93
                                                            }
                                                        }
                                                    }
                                                }
                                                switch ("" + ("\n\t" + "foo bar")) {
                                                case "\n\t":
                                                    v94 := v4[4294967295]
                                                    _ = v94
                                                case "":
                                                    var v95 uint32 = 42
                                                    _ = v95
                                                }
                                                _ = fi76
                                            }
                                            v2 = append(v2, (fi55 / (((^(fi55))) | 1)))
                                        }
                                        v96 := GenS13{F0: 1, F1: (fi55 - fi55), F2: fi55}
                                        var v97 uint8 = 127
                                        _ = v96
                                        _ = v97
                                    }
                                } else {
                                    v50 += 60
                                    fmt.Println(fi55)
                                }
                                _ = fi55
                            }
                        default:
                            var v98 uint = 2
                            _ = v98
                        }
                    }
                    var v99 int32 = 314
                    _ = v99
                } else {
                    for fi100 := 0; fi100 < 9; fi100++ {
                        v50 += -128
                        _ = fi100
                    }
                    v50 -= v50
                    if false {
                        for fi101 := 0; fi101 < 11; fi101++ {
                            v102 := v1[v48]
                            v102 = v102
                            _ = fi101
                            _ = v102
                        }
                    } else {
                        if true {
                            if true {
                                v48 = v48
                                var v103 float32 = (-2.25 / ((float32(0.0) / 1.5) - (float32(1.5) * (-2.25 + 3.14159))))
                                _ = v103
                            }
                            fmt.Println(v50, v48)
                        } else {
                            for fi104 := 0; fi104 < 9; fi104++ {
                                for fi105 := 0; fi105 < 8; fi105++ {
                                    v48 = "foo bar"
                                    _ = fi105
                                }
                                switch -1 {
                                case 9223372036854775807:
                                    var v106 string = v48
                                    _ = v106
                                case -435:
                                    var v107 string = "foo bar"
                                    _ = v107
                                case 1:
                                    var v108 uint32 = 65535
                                    _ = v108
                                }
                                _ = fi104
                            }
                            for fi109 := 0; fi109 < 18; fi109++ {
                                var v110 uint32 = 0
                                _ = fi109
                                _ = v110
                            }
                            switch (((v48 + v48) + "\n\t") + "a") {
                            case "foo bar":
                                var v111 int16 = 1
                                _ = v111
                            }
                        }
                    }
                }
            case "hello":
                v48 = "\n\t"
            case "":
                for fi112 := 0; fi112 < 2; fi112++ {
                    var v113 int32 = 42
                    for fi114 := 0; fi114 < 12; fi114++ {
                        v113 += v113
                        var v115 int = (fi112 << ((255) & 63))
                        _ = fi114
                        _ = v115
                    }
                    var v116 bool = false
                    _ = fi112
                    _ = v113
                    _ = v116
                }
            }
        } else {
            v50 += 127
            switch 1 {
            case 32767:
                v117 := GenS4{F0: ((((v48 + "\n\t") + "hello") + ((v48 + (v48 + v48)) + v48)) + v48)}
                _ = v117
            case 42:
                v118 := GenS11{F0: 0, F1: -48}
                _ = v118
            }
            for fi119 := 0; fi119 < 3; fi119++ {
                var v120 uint = 0
                v3[(fi119 * fi119)] = 2147483648
                if false {
                    switch fi119 {
                    case 32768:
                        var v121 uint8 = 2
                        _ = v121
                    case 9223372036854775807:
                        var v122 uint64 = 0
                        _ = v122
                    case 1:
                        switch ((("foo bar" + v48) + v48) + v48) {
                        case "foo bar":
                            v48 = "hello"
                        default:
                            for fi123 := 0; fi123 < 10; fi123++ {
                                switch fi119 {
                                case 228:
                                    if len(v2) > 0 {
                                        v2[0] = fi119
                                    }
                                case 42:
                                    var v124 bool = (!((v120 | (v120 << (((v120 >> (((v120 << ((65535) & 63))) & 63))) & 63))) < (v120 ^ 65535)))
                                    _ = v124
                                default:
                                    _ = len(v2)
                                }
                                _ = fi123
                            }
                        }
                    default:
                        fmt.Println(v120)
                    }
                    var v125 float64 = float64(0.0)
                    _ = v125
                } else {
                    for fi126 := 0; fi126 < 1; fi126++ {
                        v48 = (("a" + "a") + ((v48 + ("foo bar" + (v48 + v48))) + v48))
                        if false {
                            if false {
                                var v127 bool = ((v50 << (((^(v120))) & 63)) <= 1)
                                _ = v127
                            }
                            v120 = v120
                        } else {
                            var v128 float32 = 1.5
                            var v129 int16 = 248
                            _ = v128
                            _ = v129
                        }
                        if true {
                            var v130 uint64 = 2
                            v131 := GenS2{F0: int16(v130), F1: "", F2: v120}
                            fmt.Println(v48)
                            _ = v130
                            _ = v131
                        }
                        _ = fi126
                    }
                    v50 = v50
                }
                _ = fi119
                _ = v120
            }
        }
        v132 := GenS8{F0: 2147483648, F1: 57, F2: true}
        v3[-205] = 302
        _ = v132
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v53
}
