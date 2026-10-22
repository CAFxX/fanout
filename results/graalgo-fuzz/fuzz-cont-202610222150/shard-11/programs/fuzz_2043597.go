package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 int64; F1 uint64 }
type GenS11 struct { F0 uint; F1 int16; F2 int16 }
type GenS12 struct { F0 bool; F1 uint; F2 int }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 string; F1 uint8; F2 uint16 }
type GenS15 struct { F0 uint64; F1 uint16 }
type GenS16 struct { F0 bool; F1 uint16; F2 int32 }
type GenS17 struct { F0 uint }
func genFn0(p0 uint) int {
    v4 := GenS1{F0: 128}
    var v5 uint16 = 750
    if ((false && ((!(float32(1.5) < float32(0.0))) || false)) || false) {
        v5 -= 65535
        v6 := []int{861}
        _ = v6
    }
    fmt.Println(v5, p0)
    _ = p0
    _ = v4
    _ = v5
    return 32768
}

func main() {
    var v1 int = -756
    v2 := []uint32{4294967295, 1}
    var v3 int8 = 100
    v7 := GenS2{F0: v3, F1: 1}
    fmt.Println(v1)
    fmt.Println(v1)
    v3 += 127
    if true {
        var v8 uint = 26
        if false {
            if false {
                for fi9 := 0; fi9 < 6; fi9++ {
                    v2 = append(v2, 444)
                    for fi10 := 0; fi10 < 11; fi10++ {
                        fmt.Println(v8, v8)
                        v1 = v1
                        _ = fi10
                    }
                    for fi11 := 0; fi11 < 18; fi11++ {
                        if (4294967295 < 656) {
                            v2 = append(v2, 4294967295)
                        }
                        _ = fi11
                    }
                    _ = fi9
                }
                var v12 int16 = 42
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                switch v1 {
                case 128:
                    if true {
                        var v14 float32 = 3.14159
                        _ = v14
                    }
                case 0:
                    if (0.1 >= (-2.25 / 0.1)) {
                        for fi15 := 0; fi15 < 3; fi15++ {
                            fmt.Println(v1, v8, v3)
                            _ = fi15
                        }
                        var v16 uint32 = 65535
                        _ = v16
                    } else {
                        v8 = (v8 ^ v8)
                        if true {
                            _ = len(v2)
                            v3 = (v3 * (v3 | (v3 - 0)))
                        }
                    }
                case -777:
                    var v17 int64 = -661
                    _ = v17
                default:
                    switch "a" {
                    case "a":
                        var v18 uint64 = 4294967295
                        _ = v18
                    case "":
                        if false {
                            v19 := GenS3{F0: ((v3 % ((v3) | 1)) <= 74)}
                            var v20 int64 = 292
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(fi13, v1)
                            switch ("" + "hello") {
                            case "\n\t":
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = 24
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                } else {
                                    for fi21 := 0; fi21 < 4; fi21++ {
                                        var v22 int16 = -303
                                        if true {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v23 := GenS3{F0: false}
                                            fmt.Println(v8, fi13)
                                            _ = v23
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v24 := GenS2{F0: v3, F1: v8}
                                            _ = v24
                                        }
                                        v8 = (^(v8))
                                        _ = fi21
                                        _ = v22
                                    }
                                    v1 = (v1 * (^(fi13)))
                                    switch (v1 + (fi13 & v1)) {
                                    case -411:
                                        _ = len(v2)
                                    default:
                                        if (!true) {
                                            for fi25 := 0; fi25 < 11; fi25++ {
                                                if (255 <= 255) {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch fi25 {
                                                    case 2147483647:
                                                        v26 := GenS1{F0: 128}
                                                        _ = v26
                                                    case 9223372036854775807:
                                                        var v27 uint8 = 0
                                                        _ = v27
                                                    }
                                                } else {
                                                    var v28 uint64 = 18446744073709551615
                                                    switch int(v28) {
                                                    case 127:
                                                        v29 := GenS1{F0: 184}
                                                        _ = v29
                                                    case 4294967295:
                                                        var v30 uint8 = 135
                                                        _ = v30
                                                    case -575:
                                                        fmt.Println(v3, v28)
                                                    }
                                                    _ = v28
                                                }
                                                _ = fi25
                                            }
                                            var v31 uint8 = 134
                                            _ = v31
                                        } else {
                                            if false {
                                                var v32 int32 = 2147483647
                                                v8 += 18446744073709551615
                                                _ = v32
                                            } else {
                                                var v33 int8 = v3
                                                _ = v33
                                            }
                                            switch fi13 {
                                            case 42:
                                                v8 = 1
                                            case 1:
                                                switch (v1 & (v1 >> (((v8 % ((v8) | 1))) & 63))) {
                                                case 32767:
                                                    var v34 uint = (v8 % ((v8) | 1))
                                                    _ = v34
                                                case -21:
                                                    v35 := GenS2{F0: 127, F1: 216}
                                                    _ = v35
                                                default:
                                                    v36 := GenS3{F0: false}
                                                    _ = v36
                                                }
                                            case 267:
                                                v2 = append(v2, 117)
                                            default:
                                                var v37 uint16 = 65535
                                                _ = v37
                                            }
                                        }
                                    }
                                }
                            case "":
                                switch fi13 {
                                case 9223372036854775807:
                                    switch "a" {
                                    case "a":
                                        var v38 string = ""
                                        _ = v38
                                    }
                                default:
                                    v39 := GenS1{F0: 0}
                                    _ = v39
                                }
                            case "a":
                                fmt.Println(v8, v8, v1)
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    v40 := GenS4{F0: 100, F1: (v3 % ((v3) | 1))}
                                    _ = v40
                                }
                            }
                        }
                    case "\n\t":
                        if false {
                            for fi41 := 0; fi41 < 14; fi41++ {
                                v1 += (v1 * (v1 << (((v8 + (v8 - v8))) & 63)))
                                _ = fi41
                            }
                        } else {
                            var v42 uint16 = 2
                            _ = v42
                        }
                    default:
                        v2 = append(v2, 2)
                    }
                }
                _ = fi13
            }
            v1 -= v1
        }
        _ = v8
    } else {
        v1 += (-(v1))
    }
    var v43 int8 = v3
    if true {
        for fi44 := 0; fi44 < 11; fi44++ {
            if true {
                for fi45 := 0; fi45 < 13; fi45++ {
                    if ((-2.25 == 3.14159) && false) {
                        var v46 bool = (!(!(0 < 42)))
                        _ = v46
                    }
                    var v47 uint8 = 77
                    _ = len(v2)
                    _ = fi45
                    _ = v47
                }
                fmt.Println(v43, v1, v3)
            }
            _ = fi44
        }
        switch 42 {
        case 2147483647:
            v48 := GenS5{F0: 100, F1: 2, F2: 65535}
            _ = v48
        case -536:
            var v49 uint = 1
            _ = v49
        case 0:
            if ((913 >= 4294967295) && (1 <= 255)) {
                var v50 int8 = (v43 >> ((4294967295) & 63))
                v50 = v50
                v51 := GenS4{F0: (v50 | v50), F1: v3}
                _ = v50
                _ = v51
            }
        default:
            for fi52 := 0; fi52 < 11; fi52++ {
                var v53 int = (^(fi52))
                _ = fi52
                _ = v53
            }
        }
    }
    switch v1 {
    case 4294967295:
        if false {
            if (false || (!(((v1 - (v1 / ((v1) | 1))) == (^(v1))) && (581 == 255)))) {
                fmt.Println(v1, v3, v1)
                var v54 uint16 = 0
                _ = v54
            }
            if false {
                for fi55 := 0; fi55 < 18; fi55++ {
                    if (965 <= 42) {
                        if (127 != 140) {
                            var v56 int8 = v43
                            if false {
                                var v57 int = (fi55 - (fi55 * 193))
                                v1 = genFn0(758)
                                _ = v57
                            } else {
                                var v58 int32 = -990
                                _ = v58
                            }
                            var v59 bool = ("a" > "a")
                            _ = v56
                            _ = v59
                        } else {
                            var v60 int64 = 747
                            v61 := GenS4{F0: -32, F1: (v3 + (v43 + v3))}
                            _ = v60
                            _ = v61
                        }
                        v3 -= v43
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            if (true || true) {
                                v43 += v43
                                v2 = append(v2, 2147483647)
                            }
                        }
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                    }
                    _ = fi55
                }
            } else {
                switch 2147483647 {
                case -1:
                    if true {
                        fmt.Println(v43)
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v3, v43, v3)
                        case "\n\t":
                            var v62 int64 = 1
                            _ = v62
                        case "":
                            var v63 int8 = (v43 + (v43 * (v3 << ((18446744073709551615) & 63))))
                            _ = v63
                        }
                        fmt.Println(v3, v43)
                    }
                default:
                    fmt.Println(v43, v43, v3)
                }
                fmt.Println(v43, v43, v1)
            }
        } else {
            if true {
                v43 = 42
            }
            if false {
                var v64 uint = 42
                switch "" {
                case "\n\t":
                    v2 = append(v2, 2)
                case "hello":
                    var v65 uint64 = uint64(v64)
                    _ = v65
                case "foo bar":
                    if true {
                        v66 := GenS6{F0: 6, F1: (-(v1))}
                        switch (^(v1)) {
                        case -9223372036854775808:
                            if (v64 <= (^(v64))) {
                                if (!(!(!true))) {
                                    for fi67 := 0; fi67 < 20; fi67++ {
                                        for fi68 := 0; fi68 < 20; fi68++ {
                                            v69 := GenS7{F0: 2, F1: 967}
                                            var v70 bool = true
                                            _ = fi68
                                            _ = v69
                                            _ = v70
                                        }
                                        if (true || (!true)) {
                                            switch (fi67 ^ fi67) {
                                            case -1:
                                                fmt.Println(fi67, v1)
                                            case 2147483647:
                                                fmt.Println(v1, v1, fi67)
                                            }
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v71 int8 = v43
                                            for fi72 := 0; fi72 < 17; fi72++ {
                                                _ = len(v2)
                                                _ = fi72
                                            }
                                            _ = v71
                                        }
                                        var v73 uint16 = 1
                                        _ = fi67
                                        _ = v73
                                    }
                                } else {
                                    switch "a" {
                                    case "":
                                        v64 = v64
                                    case "a":
                                        v3 -= v3
                                    case "\n\t":
                                        v74 := GenS8{F0: 1}
                                        _ = v74
                                    }
                                    if (-1 == -2147483648) {
                                        var v75 int32 = -495
                                        var v76 uint32 = 463
                                        v77 := GenS9{F0: v64, F1: genFn0((v64 / ((v64) | 1)))}
                                        _ = v75
                                        _ = v76
                                        _ = v77
                                    }
                                }
                                if true {
                                    var v78 int = genFn0(v64)
                                    var v79 float32 = 0.0
                                    _ = v78
                                    _ = v79
                                }
                                fmt.Println(v3, v3, v3)
                            }
                        default:
                            if false {
                                for fi80 := 0; fi80 < 14; fi80++ {
                                    v81 := GenS10{F0: 1, F1: 1}
                                    var v82 bool = ((!true) && true)
                                    _ = fi80
                                    _ = v81
                                    _ = v82
                                }
                            }
                        }
                        _ = v66
                    } else {
                        v64 += 922
                    }
                }
                _ = v64
            } else {
                v3 = (v43 - (v3 / ((v43) | 1)))
            }
        }
    }
    if (1 != 31) {
        v3 -= v3
        if false {
            if len(v2) > 0 {
                v2[0] = 0
            }
        }
        if (1.5 != 0.1) {
            var v83 int = (v1 ^ v1)
            _ = v83
        } else {
            v1 -= (v1 % ((v1) | 1))
            var v84 uint64 = 42
            _ = v84
        }
    } else {
        switch (v1 / ((v1) | 1)) {
        case -9223372036854775808:
            var v85 int32 = 2147483647
            _ = v85
        }
        if true {
            v3 = v3
        } else {
            var v86 uint64 = 2
            _ = v86
        }
    }
    if false {
        fmt.Println(v3, v1, v1)
        fmt.Println(v43)
        if true {
            var v87 float64 = (1e10 * float64((1.5 - 1e10)))
            v88 := GenS11{F0: 519, F1: 0, F2: 0}
            _ = v87
            _ = v88
        } else {
            v89 := GenS12{F0: (!false), F1: 42, F2: v1}
            _ = v89
        }
    } else {
        if (1 >= 981) {
            v90 := GenS7{F0: 0, F1: 4294967295}
            var v91 uint32 = 4294967295
            var v92 uint16 = 42
            _ = v90
            _ = v91
            _ = v92
        }
        v43 += v43
        v1 -= 128
    }
    switch (v1 << ((18446744073709551615) & 63)) {
    case 42:
        var v93 int64 = 34
        _ = v93
    case 311:
        v3 += 55
    case -9223372036854775808:
        switch v1 {
        case 373:
            var v94 float64 = (((float64(-2.25) * 0.0) + (1.5 / 0.1)) + 3.14159)
            _ = v94
        case 128:
            var v95 int = (v1 % (((v1 + genFn0(0))) | 1))
            _ = v95
        default:
            if true {
                v43 = v3
                for fi96 := 0; fi96 < 10; fi96++ {
                    var v97 uint16 = 293
                    _ = fi96
                    _ = v97
                }
                for fi98 := 0; fi98 < 14; fi98++ {
                    if false {
                        for fi99 := 0; fi99 < 14; fi99++ {
                            v1 += (-(fi98))
                            fmt.Println(fi98)
                            _ = fi99
                        }
                        v3 += v43
                        var v100 int16 = -888
                        _ = v100
                    }
                    if true {
                        switch "a" {
                        case "\n\t":
                            v101 := GenS4{F0: 0, F1: (v3 * (v3 | (v3 | 42)))}
                            _ = v101
                        case "foo bar":
                            var v102 int32 = 0
                            _ = v102
                        }
                        v103 := GenS5{F0: v43, F1: 0, F2: 65535}
                        fmt.Println(fi98)
                        _ = v103
                    } else {
                        v2 = append(v2, 109)
                        v1 = fi98
                    }
                    _ = fi98
                }
            } else {
                v104 := GenS13{F0: 42}
                if true {
                    var v105 int16 = 1
                    _ = v105
                } else {
                    v1 += (v1 >> ((372) & 63))
                }
                fmt.Println(v1, v1, v3)
                _ = v104
            }
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((float64((0.1 + (0.0 - 0.0))) - 0.0) > 1e10) {
        for fi106 := 0; fi106 < 1; fi106++ {
            v43 -= v43
            v1 = fi106
            if true {
                if ((32767 <= 2) || true) {
                    var v107 uint32 = 1
                    if len(v2) > 0 {
                        v2[0] = (v107 % ((v107) | 1))
                    }
                    fmt.Println(v107, v1)
                    _ = v107
                } else {
                    var v108 string = "hello"
                    _ = v108
                }
                fmt.Println(v1)
            } else {
                if (42 <= 18446744073709551615) {
                    var v109 string = "\n\t"
                    fmt.Println(v43)
                    _ = v109
                } else {
                    if ("\n\t" != (("hello" + "\n\t") + "foo bar")) {
                        v1 -= (fi106 & (fi106 + v1))
                        v110 := GenS10{F0: -1, F1: 846}
                        _ = v110
                    }
                    for fi111 := 0; fi111 < 13; fi111++ {
                        var v112 bool = true
                        fmt.Println(v43, v43)
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                        var v113 uint = 253
                        var v114 uint64 = 485
                        var v115 uint32 = 65535
                        v114 = 18446744073709551615
                        if (v112 || false) {
                            var v116 int8 = 42
                            var v117 uint = v113
                            _ = v116
                            _ = v117
                        } else {
                            fmt.Println(v113)
                        }
                        _ = fi111
                        _ = v112
                        _ = v113
                        _ = v114
                        _ = v115
                    }
                    v118 := GenS14{F0: (("" + "") + ""), F1: 56, F2: 32767}
                    _ = v118
                }
                if (false && false) {
                    fmt.Println(fi106, v3, v3)
                }
                v3 = v3
            }
            _ = fi106
        }
        v119 := GenS15{F0: 118, F1: 2}
        for fi120 := 0; fi120 < 5; fi120++ {
            for fi121 := 0; fi121 < 13; fi121++ {
                v3 = v3
                _ = fi121
            }
            for fi122 := 0; fi122 < 18; fi122++ {
                var v123 int32 = -199
                fmt.Println(v43, v43, fi122)
                if true {
                    v123 -= (v123 * 2147483647)
                }
                _ = fi122
                _ = v123
            }
            v3 = (-(v3))
            _ = fi120
        }
        _ = v119
    } else {
        v124 := GenS1{F0: 128}
        _ = v124
    }
    for fi125 := 0; fi125 < 8; fi125++ {
        var v126 uint8 = 1
        var v127 int64 = 42
        var v128 uint64 = 13
        _ = fi125
        _ = v126
        _ = v127
        _ = v128
    }
    for fi129 := 0; fi129 < 11; fi129++ {
        if ((!(false && false)) && false) {
            var v130 int = 188
            if (1.5 > (1e10 + 1.5)) {
                v3 += v43
                for fi131 := 0; fi131 < 12; fi131++ {
                    v132 := GenS16{F0: ("" <= "foo bar"), F1: 65535, F2: 2147483647}
                    v130 -= v1
                    _ = fi131
                    _ = v132
                }
                var v133 int = v1
                _ = v133
            }
            _ = v130
        } else {
            var v134 int16 = 556
            _ = v134
        }
        var v135 uint8 = 146
        _ = fi129
        _ = v135
    }
    fmt.Println(v1, v1)
    for fi136 := 0; fi136 < 5; fi136++ {
        v137 := GenS12{F0: (((255 > 80) && true) && true), F1: 255, F2: fi136}
        v138 := GenS17{F0: 18446744073709551615}
        v2 = append(v2, 809)
        _ = fi136
        _ = v137
        _ = v138
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v7
}
