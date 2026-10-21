package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint; F1 int32; F2 int64 }
type GenS7 struct { F0 uint16; F1 uint64; F2 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool; F1 uint8; F2 uint }
type GenS10 struct { F0 string; F1 uint16; F2 int16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int64; F1 uint8; F2 uint16 }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int = (-(fi6))
        fmt.Println(fi6, v7)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 29, F1: 9223372036854775807, F2: 214}
    _ = v8
    return 12
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = v1
    var v3 int8 = -43
    var v4 uint32 = 2147483647
    var v5 int64 = 9223372036854775807
    switch 1 {
    case -2147483648:
        v9 := GenS1{F0: 127, F1: -883, F2: 0}
        _ = v9
    case 32767:
        fmt.Println(v1)
    case -577:
        switch "hello" {
        case "\n\t":
            v4 -= v4
        case "":
            v2 = -839
        case "a":
            v10 := []int{32767, int(v5), -1, 4294967295}
            _ = v10
        }
    }
    v5 -= -71
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := []int{(fi11 ^ fi11), fi11}
        _ = len(v12)
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: 254, F1: fi13, F2: 65535}
        v15 := []int{-1, (fi13 % (((fi13 & -301)) | 1)), 0}
        if len(v15) > 0 {
            v15[0] = (fi13 / ((-9223372036854775808) | 1))
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 42}
    v17 := []int{127, -9223372036854775808, 204}
    for fi18 := 0; fi18 < 10; fi18++ {
        if false {
            for fi19 := 0; fi19 < 2; fi19++ {
                fmt.Println(v4, fi18)
                v20 := GenS2{F0: (v5 / ((-60) | 1))}
                if true {
                    v5 -= (v5 * v5)
                    var v21 uint = uint(v4)
                    _ = v21
                } else {
                    v17 = append(v17, fi18)
                    v22 := GenS2{F0: -591}
                    var v23 int64 = v5
                    _ = v22
                    _ = v23
                }
                _ = fi19
                _ = v20
            }
            v5 = v5
        } else {
            var v24 int32 = (v1 * (v1 % (((v1 << ((2) & 63))) | 1)))
            switch ("" + "hello") {
            case "foo bar":
                for fi25 := 0; fi25 < 11; fi25++ {
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v5)
                    case "hello":
                        var v26 float64 = 0.1
                        _ = v26
                    case "":
                        if true {
                            switch fi25 {
                            case 4294967295:
                                switch "" {
                                case "":
                                    v24 = -173
                                case "\n\t":
                                    if false {
                                        v27 := GenS1{F0: 2, F1: (fi18 & (fi18 >> ((42) & 63))), F2: 2}
                                        _ = v27
                                    } else {
                                        var v28 float32 = 1.5
                                        _ = v28
                                    }
                                default:
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(v4)
                                    default:
                                        v24 += (v24 >> ((18446744073709551615) & 63))
                                    }
                                }
                            case 567:
                                var v29 float64 = (0.0 + (-2.25 + 1.5))
                                _ = v29
                            }
                            var v30 string = (("hello" + "") + ("a" + "\n\t"))
                            if len(v17) > 0 {
                                v17[0] = -84
                            }
                            _ = v30
                        } else {
                            v31 := GenS2{F0: (-(v5))}
                            _ = v31
                        }
                    default:
                        fmt.Println(fi25)
                    }
                    _ = fi25
                }
            case "":
                for fi32 := 0; fi32 < 15; fi32++ {
                    if true {
                        fmt.Println(fi32, v5, fi18)
                        v33 := GenS3{F0: 1}
                        if len(v17) > 0 {
                            v17[0] = fi18
                        }
                        _ = v33
                    }
                    if (!true) {
                        if (((("hello" + (("\n\t" + ("foo bar" + "a")) + ("foo bar" + ("foo bar" + "hello")))) + "a") + "foo bar") < "hello") {
                            v17 = append(v17, fi32)
                            if true {
                                _ = len(v17)
                                if len(v17) > 0 {
                                    v17[0] = 4294967295
                                }
                                fmt.Println(fi32)
                            }
                        }
                        v3 += (^(v3))
                    } else {
                        if len(v17) > 0 {
                            v17[0] = (fi32 % ((fi18) | 1))
                        }
                    }
                    _ = fi32
                }
            case "\n\t":
                switch "hello" {
                case "foo bar":
                    for fi34 := 0; fi34 < 10; fi34++ {
                        var v35 int16 = -32768
                        var v36 int16 = -588
                        v5 -= (v5 % (((^(v5))) | 1))
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                case "hello":
                    if (true || true) {
                        _ = len(v17)
                        if false {
                            for fi37 := 0; fi37 < 9; fi37++ {
                                v4 = v4
                                _ = fi37
                            }
                            switch fi18 {
                            case 559:
                                v5 += v5
                            case -1:
                                for fi38 := 0; fi38 < 11; fi38++ {
                                    var v39 uint16 = 543
                                    v2 -= (v2 - -1)
                                    var v40 uint32 = (v4 % ((uint32(v39)) | 1))
                                    switch fi38 {
                                    case 4294967295:
                                        v41 := GenS1{F0: 229, F1: 32768, F2: uint16(v40)}
                                        _ = v41
                                    case -9223372036854775808:
                                        v4 -= v4
                                    default:
                                        var v42 uint64 = 18446744073709551615
                                        _ = v42
                                    }
                                    switch int(v24) {
                                    case -8:
                                        if true {
                                            fmt.Println(v40, v2, v5)
                                            v43 := GenS4{F0: 1, F1: (fi18 * 127)}
                                            _ = v43
                                        }
                                    case 2147483647:
                                        switch (fi18 / ((fi18) | 1)) {
                                        case -2147483648:
                                            var v44 float32 = float32(1e10)
                                            _ = v44
                                        }
                                    }
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                            default:
                                v45 := GenS1{F0: 225, F1: fi18, F2: 431}
                                _ = v45
                            }
                            v46 := GenS1{F0: 1, F1: fi18, F2: 2}
                            _ = v46
                        } else {
                            v1 = v1
                            v47 := GenS2{F0: (v5 ^ v5)}
                            for fi48 := 0; fi48 < 20; fi48++ {
                                var v49 int32 = (v1 & (v1 << ((1) & 63)))
                                var v50 uint32 = (v4 * v4)
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            _ = v47
                        }
                        if true {
                            for fi51 := 0; fi51 < 8; fi51++ {
                                v52 := GenS3{F0: uint8(v4)}
                                _ = fi51
                                _ = v52
                            }
                        } else {
                            fmt.Println(v5, v5, fi18)
                            var v53 uint32 = v4
                            _ = v53
                        }
                    }
                case "":
                    var v54 int32 = v2
                    _ = v54
                default:
                    var v55 uint = uint(v4)
                    _ = v55
                }
            }
            if false {
                var v56 uint64 = 817
                var v57 int32 = int32(v5)
                _ = v56
                _ = v57
            }
            _ = v24
        }
        _ = fi18
    }
    for fi58 := 0; fi58 < 19; fi58++ {
        fmt.Println(v4, v1)
        v4 = v4
        _ = fi58
    }
    if true {
        var v59 int = -1
        _ = v59
    }
    v2 += (v2 << ((255) & 63))
    if false {
        for fi60 := 0; fi60 < 15; fi60++ {
            v61 := GenS5{F0: "foo bar", F1: (1000 >= int16(fi60))}
            _ = fi60
            _ = v61
        }
        for fi62 := 0; fi62 < 7; fi62++ {
            v1 = v2
            _ = fi62
        }
    } else {
        v63 := GenS2{F0: int64(v4)}
        switch 251 {
        case 127:
            v17 = append(v17, 873)
        case -835:
            var v64 uint = 318
            _ = v64
        default:
            switch ("hello" + "hello") {
            case "hello":
                v65 := GenS3{F0: 123}
                _ = v65
            case "foo bar":
                if false {
                    for fi66 := 0; fi66 < 17; fi66++ {
                        for fi67 := 0; fi67 < 19; fi67++ {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            switch (fi67 + fi66) {
                            case -981:
                                if (!false) {
                                    switch (("foo bar" + "foo bar") + "\n\t") {
                                    case "foo bar":
                                        switch fi66 {
                                        case -922:
                                            var v68 float64 = (0.1 / (((1e10 * 0.0) - (((1e10 * float64((3.14159 + 3.14159))) * 3.14159) + (float64(0.0) - 0.1))) / float64(3.14159)))
                                            _ = v68
                                        case 139:
                                            var v69 float64 = 1.5
                                            _ = v69
                                        default:
                                            if true {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    fmt.Println(v3)
                                                case "":
                                                    switch 668 {
                                                    case 0:
                                                        var v70 int16 = -209
                                                        _ = v70
                                                    }
                                                case "\n\t":
                                                    v5 -= (v5 ^ v5)
                                                }
                                                v17 = append(v17, fi66)
                                            } else {
                                                v5 += -1
                                            }
                                        }
                                    case "\n\t":
                                        var v71 uint = 2
                                        _ = v71
                                    case "":
                                        v5 -= (v5 & v5)
                                    default:
                                        if (0 >= 127) {
                                            switch fi67 {
                                            case -790:
                                                switch fi67 {
                                                case -783:
                                                    fmt.Println(v4, v5, v3)
                                                case 9223372036854775807:
                                                    var v72 uint64 = 18446744073709551615
                                                    _ = v72
                                                case 4294967295:
                                                    v73 := GenS4{F0: v1, F1: (-(fi66))}
                                                    _ = v73
                                                }
                                            case 321:
                                                var v74 uint8 = 176
                                                _ = v74
                                            }
                                        } else {
                                            v75 := GenS1{F0: 124, F1: -35, F2: 770}
                                            var v76 uint = 65535
                                            _ = v75
                                            _ = v76
                                        }
                                        var v77 float64 = float64(1.5)
                                        switch ("a" + "foo bar") {
                                        case "foo bar":
                                            v78 := GenS6{F0: 255, F1: v2, F2: (-(v5))}
                                            _ = v78
                                        }
                                        _ = v77
                                    }
                                    if false {
                                        var v79 float64 = 1e10
                                        _ = v79
                                    }
                                    if (v5 != 42) {
                                        var v80 int64 = v5
                                        switch (fi66 % ((fi66) | 1)) {
                                        case 2147483647:
                                            var v81 float32 = 0.0
                                            _ = v81
                                        default:
                                            var v82 uint16 = 0
                                            _ = v82
                                        }
                                        _ = v80
                                    } else {
                                        v83 := GenS7{F0: 32767, F1: 1, F2: 128}
                                        _ = v83
                                    }
                                } else {
                                    v4 += (v4 * (v4 + (v4 << ((42) & 63))))
                                    v3 = v3
                                    if true {
                                        v1 += v2
                                        if ((v3 - v3) < v3) {
                                            v84 := GenS8{F0: (^(fi67))}
                                            _ = v84
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                fmt.Println(fi66, fi66, fi66)
                                            case "a":
                                                v4 -= v4
                                                var v85 uint16 = 2
                                                _ = v85
                                            }
                                            if true {
                                                if (true || (42 != 42)) {
                                                    v4 = v4
                                                    if len(v17) >= 2 {
                                                        v17 = v17[1:len(v17)]
                                                    }
                                                    if len(v17) > 0 {
                                                        v17[0] = int(v5)
                                                    }
                                                } else {
                                                    var v86 bool = (int16(v3) > 50)
                                                    if false {
                                                        if v86 {
                                                            var v87 bool = v86
                                                            fmt.Println(v1, v4)
                                                            fmt.Println(fi66)
                                                            _ = v87
                                                        } else {
                                                            if len(v17) >= 2 {
                                                                v17 = v17[1:len(v17)]
                                                            }
                                                            if (!v86) {
                                                                var v88 uint = 18446744073709551615
                                                                v89 := GenS7{F0: 32767, F1: 0, F2: (^(fi67))}
                                                                _ = v88
                                                                _ = v89
                                                            } else {
                                                                v4 += 65535
                                                            }
                                                            if (!(!v86)) {
                                                                if v86 {
                                                                    v2 -= v2
                                                                    var v90 int16 = 0
                                                                    _ = v90
                                                                } else {
                                                                    v2 += int32(v3)
                                                                    fmt.Println(fi67, v2)
                                                                    var v91 float64 = ((1.5 + 0.0) * 1e10)
                                                                    _ = v91
                                                                }
                                                            }
                                                            v17 = append(v17, fi66)
                                                        }
                                                    }
                                                    if v86 {
                                                        var v92 bool = v86
                                                        _ = v92
                                                    }
                                                    _ = v86
                                                }
                                                if false {
                                                    var v93 uint8 = 2
                                                    _ = v93
                                                }
                                                v17 = append(v17, fi66)
                                            } else {
                                                v5 = (v5 * v5)
                                                v94 := GenS9{F0: false, F1: 128, F2: 18446744073709551615}
                                                _ = v94
                                            }
                                            v95 := GenS2{F0: 203}
                                            _ = v95
                                        }
                                    }
                                    var v96 int = (fi66 | fi67)
                                    var v97 int8 = v3
                                    var v98 uint8 = 81
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    _ = v96
                                    _ = v97
                                    _ = v98
                                }
                            case -970:
                                switch -2147483648 {
                                case -1:
                                    var v99 uint16 = 1
                                    _ = v99
                                case 32767:
                                    switch "\n\t" {
                                    case "foo bar":
                                        if (((true && true) && false) || ((1e10 / 1.0) >= float32((0.1 - (((0.0 / 0.1) + 3.14159) * 1.5))))) {
                                            v5 = v5
                                            v17 = append(v17, (fi66 >> ((42) & 63)))
                                        } else {
                                            fmt.Println(v1, v2)
                                            switch "\n\t" {
                                            case "":
                                                var v100 float64 = 1e10
                                                _ = v100
                                            default:
                                                v1 = 2147483647
                                            }
                                            v1 += (v2 ^ (v2 + v1))
                                        }
                                        var v101 uint8 = 194
                                        _ = v101
                                    }
                                default:
                                    if false {
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                        var v102 string = "foo bar"
                                        switch v102 {
                                        case "":
                                            var v103 float32 = 0.0
                                            _ = v103
                                        case "\n\t":
                                            v3 = -122
                                        case "a":
                                            var v104 uint16 = 1
                                            _ = v104
                                        default:
                                            var v105 uint8 = 97
                                            _ = v105
                                        }
                                        _ = v102
                                    }
                                }
                            }
                            _ = fi67
                        }
                        v5 += v5
                        _ = fi66
                    }
                } else {
                    if false {
                        for fi106 := 0; fi106 < 17; fi106++ {
                            var v107 bool = false
                            for fi108 := 0; fi108 < 17; fi108++ {
                                v17 = append(v17, fi106)
                                if v107 {
                                    switch "a" {
                                    case "":
                                        v109 := GenS9{F0: false, F1: 127, F2: 4294967295}
                                        _ = v109
                                    }
                                    v1 = 0
                                    switch "" {
                                    case "":
                                        v110 := GenS10{F0: "\n\t", F1: 599, F2: -32768}
                                        v1 = v2
                                        v5 += 1
                                        v107 = ((-(v5)) >= (-(v5)))
                                        _ = v110
                                    case "hello":
                                        if (!(42 == 42)) {
                                            switch fi108 {
                                            case -36:
                                                if (v1 != 711) {
                                                    if len(v17) > 0 {
                                                        v17[0] = -616
                                                    }
                                                } else {
                                                    var v111 bool = false
                                                    var v112 bool = false
                                                    _ = v111
                                                    _ = v112
                                                }
                                            case 127:
                                                _ = len(v17)
                                            default:
                                                if v107 {
                                                    if ((368 >= v4) && v107) {
                                                        if v107 {
                                                            switch (fi108 | (fi106 << ((0) & 63))) {
                                                            case 400:
                                                                var v113 uint = 65535
                                                                _ = v113
                                                            default:
                                                                if v107 {
                                                                    var v114 uint = 283
                                                                    if v107 {
                                                                        if len(v17) > 0 {
                                                                            v17[0] = int(v3)
                                                                        }
                                                                        if v107 {
                                                                            v5 = v5
                                                                        } else {
                                                                            if v107 {
                                                                                fmt.Println(fi108, v2, v114)
                                                                            } else {
                                                                                var v115 uint = 2
                                                                                var v116 uint = 111
                                                                                _ = v115
                                                                                _ = v116
                                                                            }
                                                                        }
                                                                        v1 -= v2
                                                                    } else {
                                                                        if v107 {
                                                                            if len(v17) > 0 {
                                                                                v17[0] = (fi108 % (((fi108 & (-(fi108)))) | 1))
                                                                            }
                                                                            v5 += v5
                                                                        } else {
                                                                            var v117 string = ("hello" + ("\n\t" + (("\n\t" + ("\n\t" + "foo bar")) + (("a" + "foo bar") + ""))))
                                                                            switch fi108 {
                                                                            case 42:
                                                                                var v118 float32 = 1e10
                                                                                _ = v118
                                                                            }
                                                                            _ = v117
                                                                        }
                                                                        v4 += (v4 << (((v114 << (((v114 | 0)) & 63))) & 63))
                                                                        if len(v17) >= 2 {
                                                                            v17 = v17[1:len(v17)]
                                                                        }
                                                                        v4 -= 1
                                                                    }
                                                                    _ = v114
                                                                }
                                                            }
                                                            var v119 int32 = v2
                                                            _ = v119
                                                        } else {
                                                            v120 := GenS8{F0: (fi106 << ((976) & 63))}
                                                            _ = len(v17)
                                                            if v107 {
                                                                switch 4294967295 {
                                                                case -9223372036854775808:
                                                                    if len(v17) > 0 {
                                                                        v17[0] = 32768
                                                                    }
                                                                }
                                                                var v121 int32 = -1
                                                                _ = v121
                                                            }
                                                            _ = v120
                                                        }
                                                        v122 := GenS11{F0: (v4 >> ((109) & 63))}
                                                        if (!(1 < 218)) {
                                                            if true {
                                                                if v107 {
                                                                    var v123 uint = 0
                                                                    var v124 uint16 = 65535
                                                                    _ = v123
                                                                    _ = v124
                                                                }
                                                                var v125 bool = (v107 && (fi108 > fi108))
                                                                _ = v125
                                                            } else {
                                                                var v126 uint = 215
                                                                _ = v126
                                                            }
                                                            v127 := GenS12{F0: 42}
                                                            _ = v127
                                                        } else {
                                                            var v128 uint64 = 2
                                                            if v107 {
                                                                var v129 string = "a"
                                                                _ = v129
                                                            } else {
                                                                v130 := GenS10{F0: "a", F1: 799, F2: 32767}
                                                                fmt.Println(v2)
                                                                _ = v130
                                                            }
                                                            _ = len(v17)
                                                            _ = v128
                                                        }
                                                        _ = v122
                                                    } else {
                                                        fmt.Println(v3, v107)
                                                    }
                                                    if (42 != 42) {
                                                        if (true && (!(!v107))) {
                                                            v131 := GenS8{F0: fi106}
                                                            v132 := GenS9{F0: (0 < 2), F1: 46, F2: 615}
                                                            fmt.Println(fi106, v4, fi108)
                                                            _ = v131
                                                            _ = v132
                                                        } else {
                                                            var v133 int8 = 127
                                                            _ = v133
                                                        }
                                                    }
                                                    switch ("a" + (((("foo bar" + "foo bar") + (("foo bar" + "") + "foo bar")) + "hello") + "foo bar")) {
                                                    case "foo bar":
                                                        fmt.Println(v1, v3)
                                                    case "hello":
                                                        v134 := GenS4{F0: v1, F1: (fi106 % ((fi108) | 1))}
                                                        _ = v134
                                                    }
                                                    var v135 int64 = v5
                                                    _ = v135
                                                }
                                            }
                                            var v136 float64 = float64(1.5)
                                            _ = v136
                                        }
                                    case "\n\t":
                                        if v107 {
                                            if ("foo bar" < "foo bar") {
                                                fmt.Println(v5)
                                                v137 := GenS13{F0: 2}
                                                if true {
                                                    var v138 string = (("\n\t" + "hello") + "\n\t")
                                                    var v139 int = (fi108 | (fi106 << ((255) & 63)))
                                                    if v107 {
                                                        v140 := GenS14{F0: (v5 >> ((420) & 63)), F1: 2, F2: 65535}
                                                        _ = v140
                                                    }
                                                    v17 = append(v17, -9223372036854775808)
                                                    if len(v17) >= 2 {
                                                        v17 = v17[1:len(v17)]
                                                    }
                                                    _ = v138
                                                    _ = v139
                                                } else {
                                                    var v141 string = "foo bar"
                                                    v107 = (v107 && v107)
                                                    var v142 int16 = -864
                                                    _ = v141
                                                    _ = v142
                                                }
                                                _ = v137
                                            }
                                            v4 = v4
                                            var v143 int = fi108
                                            v144 := GenS2{F0: (^(v5))}
                                            fmt.Println(v2, v1)
                                            var v145 uint64 = genFn0()
                                            switch (fi106 & (v143 + 2147483647)) {
                                            case 127:
                                                var v146 string = (("\n\t" + "foo bar") + "foo bar")
                                                var v147 string = v146
                                                var v148 int = fi106
                                                _ = v146
                                                _ = v147
                                                _ = v148
                                            case -298:
                                                v107 = (!(!false))
                                            }
                                            _ = v143
                                            _ = v144
                                            _ = v145
                                        } else {
                                            var v149 int16 = 0
                                            var v150 bool = v107
                                            _ = v149
                                            _ = v150
                                        }
                                    }
                                } else {
                                    fmt.Println(v107)
                                    var v151 string = "hello"
                                    _ = v151
                                }
                                _ = fi108
                            }
                            v152 := GenS10{F0: ("foo bar" + (("a" + "foo bar") + ("\n\t" + ""))), F1: 32767, F2: -444}
                            _ = fi106
                            _ = v107
                            _ = v152
                        }
                    }
                }
            case "\n\t":
                for fi153 := 0; fi153 < 19; fi153++ {
                    v2 = int32(fi153)
                    if len(v17) > 0 {
                        v17[0] = (fi153 / ((999) | 1))
                    }
                    fmt.Println(v5)
                    _ = fi153
                }
            }
        }
        _ = v63
    }
    v2 = v2
    v4 -= v4
    v5 = 0
    var v154 float64 = float64((3.14159 * 0.1))
    var v155 int16 = 1
    if ((!(!(v2 < -881))) && (255 < 2)) {
        if false {
            v3 -= 42
            for fi156 := 0; fi156 < 3; fi156++ {
                var v157 int = fi156
                fmt.Println(v2)
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = fi156
                _ = v157
            }
        }
    } else {
        var v158 uint32 = (^(v4))
        _ = v158
    }
    v154 += (v154 * (float64((0.1 * (0.0 - float32(v154)))) + v154))
    fmt.Println("v1:", v1)
    fmt.Println("v154:", v154)
    fmt.Println("v155:", v155)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v154
    _ = v155
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
