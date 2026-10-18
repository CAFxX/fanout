package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 int; F2 int8 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 string; F1 uint }
type GenS9 struct { F0 uint64; F1 uint8 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 bool; F1 int8 }
type GenS13 struct { F0 string; F1 int16; F2 bool }
type GenS14 struct { F0 uint8; F1 uint16; F2 string }
type GenS15 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint = 299
    v2 := []int{42, 556}
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]int)
    v5 := []int{32767, -1, -214}
    v6 := []int{int(v1), -1}
    if (127 <= 0) {
        var v7 uint8 = 1
        for fi8 := 0; fi8 < 8; fi8++ {
            switch fi8 {
            case -2147483648:
                if (!("foo bar" != "hello")) {
                    var v9 float64 = ((1e10 / (float64(((1.5 / 0.1) - -2.25)) * (1.5 * ((0.0 - 0.1) - float64((-2.25 + float32(0.0))))))) / float64(3.14159))
                    _ = len(v3)
                    if true {
                        delete(v4, 464)
                        for fi10 := 0; fi10 < 2; fi10++ {
                            fmt.Println(v7)
                            v1 += (v1 << ((v1) & 63))
                            _ = fi10
                        }
                        for fi11 := 0; fi11 < 15; fi11++ {
                            var v12 float32 = 0.0
                            v13 := GenS1{F0: v7}
                            var v14 uint16 = 0
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    } else {
                        var v15 int = fi8
                        v9 = (((0.0 - v9) * ((v9 / 0.1) + 1e10)) * v9)
                        fmt.Println(fi8, v7)
                        _ = v15
                    }
                    _ = v9
                }
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 uint32 = 176
                var v18 uint64 = 833
                switch (fi16 / (((fi8 * (fi16 & (fi16 + -2147483648)))) | 1)) {
                case 612:
                    var v19 int16 = -1
                    _ = v19
                case 586:
                    fmt.Println(fi16)
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            var v20 uint16 = 2
            _ = fi8
            _ = v20
        }
        _ = v7
    } else {
        if (true && ((255 != v1) || (((0.1 / (0.0 - 0.1)) <= 0.1) || (!false)))) {
            for fi21 := 0; fi21 < 4; fi21++ {
                fmt.Println(fi21, fi21)
                _ = fi21
            }
            v3[428] = false
            delete(v4, 501)
        } else {
            v1 += v1
            v22 := GenS1{F0: uint8(v1)}
            var v23 uint32 = 65535
            _ = v22
            _ = v23
        }
        var v24 uint = (v1 ^ v1)
        var v25 int32 = 2147483647
        _ = v24
        _ = v25
    }
    if true {
        var v26 uint8 = 160
        var v27 int = -290
        var v28 float64 = 3.14159
        _ = v26
        _ = v27
        _ = v28
    } else {
        for fi29 := 0; fi29 < 3; fi29++ {
            for fi30 := 0; fi30 < 2; fi30++ {
                delete(v4, 1)
                var v31 uint16 = 32767
                _ = fi30
                _ = v31
            }
            v1 = (v1 + (v1 << ((42) & 63)))
            fmt.Println(v1, fi29)
            _ = fi29
        }
        v1 = (v1 - v1)
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        if ((("hello" + "") + ((("" + "a") + ("" + "\n\t")) + ("a" + "a"))) >= "hello") {
            fmt.Println(v1, fi32)
            var v33 uint8 = 3
            _ = v33
        } else {
            v34 := GenS2{F0: 42, F1: v1, F2: 183}
            _ = v34
        }
        _ = fi32
    }
    v1 -= 827
    v1 -= (v1 >> ((v1) & 63))
    for fi35 := 0; fi35 < 4; fi35++ {
        if (false && true) {
            v1 = 616
            v1 += (v1 + (v1 | (v1 | v1)))
        } else {
            v1 -= (v1 >> ((v1) & 63))
            for fi36 := 0; fi36 < 14; fi36++ {
                v1 -= (v1 << ((0) & 63))
                _ = fi36
            }
            var v37 bool = (((3.14159 * (3.14159 * 3.14159)) - 0.1) == float64((0.0 / 0.1)))
            _ = v37
        }
        if false {
            for fi38 := 0; fi38 < 16; fi38++ {
                if false {
                    v1 = 730
                    v1 += 258
                    v1 = 65535
                }
                _ = fi38
            }
        } else {
            var v39 int = 548
            switch ("hello" + ("" + "foo bar")) {
            case "":
                if false {
                    if ((0 == 2147483647) || true) {
                        fmt.Println(fi35, fi35, fi35)
                        v39 += 0
                        var v40 int = 128
                        _ = v40
                    }
                    if true {
                        for fi41 := 0; fi41 < 19; fi41++ {
                            if true {
                                v1 -= v1
                                var v42 int32 = 36
                                _ = v42
                            } else {
                                var v43 uint64 = 42
                                v43 = v43
                                _ = v43
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v44 int = fi35
                            _ = fi41
                            _ = v44
                        }
                        if true {
                            v39 += int(v1)
                            if true {
                                v45 := GenS1{F0: 95}
                                for fi46 := 0; fi46 < 8; fi46++ {
                                    v39 = fi46
                                    _ = fi46
                                }
                                _ = v45
                            } else {
                                switch (v39 << (((v1 - 949)) & 63)) {
                                case -29:
                                    var v47 bool = true
                                    _ = v47
                                case 32767:
                                    fmt.Println(fi35, v39, v1)
                                default:
                                    var v48 uint16 = 2
                                    _ = v48
                                }
                                var v49 float32 = ((0.0 / (-2.25 / (1e10 / (3.14159 + 0.0)))) / (1.5 + -2.25))
                                switch (v39 * int(v1)) {
                                case -112:
                                    v50 := GenS3{F0: -9223372036854775808, F1: 1}
                                    _ = v50
                                case 2147483647:
                                    for fi51 := 0; fi51 < 13; fi51++ {
                                        var v52 uint8 = 128
                                        v53 := GenS4{F0: 904, F1: 65535}
                                        if false {
                                            v1 = 1
                                        } else {
                                            v1 += v1
                                            fmt.Println(v1)
                                        }
                                        _ = fi51
                                        _ = v52
                                        _ = v53
                                    }
                                }
                                _ = v49
                            }
                        } else {
                            var v54 int64 = 0
                            for fi55 := 0; fi55 < 9; fi55++ {
                                var v56 string = "\n\t"
                                var v57 uint8 = 1
                                _ = fi55
                                _ = v56
                                _ = v57
                            }
                            v58 := GenS2{F0: 2, F1: v1, F2: 255}
                            _ = v54
                            _ = v58
                        }
                        v59 := GenS2{F0: uint64(v1), F1: v1, F2: 2}
                        _ = v59
                    } else {
                        v60 := GenS5{F0: 100}
                        _ = v60
                    }
                    switch "hello" {
                    case "":
                        v39 += (v39 / ((v39) | 1))
                    default:
                        switch v39 {
                        case 32767:
                            for fi61 := 0; fi61 < 19; fi61++ {
                                if false {
                                    var v62 uint = v1
                                    _ = v62
                                }
                                if true {
                                    switch (fi35 % ((fi61) | 1)) {
                                    case 279:
                                        fmt.Println(v1, v1)
                                    default:
                                        v39 = fi35
                                    }
                                } else {
                                    var v63 int32 = 2147483647
                                    var v64 uint32 = 65535
                                    fmt.Println(v39, fi61)
                                    switch "hello" {
                                    case "foo bar":
                                        if false {
                                            var v65 int = -2147483648
                                            _ = v65
                                        } else {
                                            if false {
                                                if false {
                                                    if len(v2) > 0 {
                                                        v2[0] = (v39 << (((^(v1))) & 63))
                                                    }
                                                    v3[259] = false
                                                } else {
                                                    v39 += (v39 & -333)
                                                    v39 -= -2147483648
                                                    fmt.Println(v39, fi61)
                                                }
                                            }
                                            v39 = (v39 / (((fi35 % (((-(fi35))) | 1))) | 1))
                                        }
                                    case "hello":
                                        v1 = (^(v1))
                                    }
                                    fmt.Println(v39, v39)
                                    v66 := GenS6{F0: ("" + "foo bar"), F1: (fi61 - (fi35 | v39)), F2: -1}
                                    fmt.Println(fi35, fi35, fi61)
                                    var v67 int8 = 42
                                    if true {
                                        v68 := GenS6{F0: "\n\t", F1: (^(fi35)), F2: v67}
                                        if true {
                                            if true {
                                                if (1 == 60) {
                                                    v69 := GenS1{F0: 255}
                                                    switch v39 {
                                                    case 2147483647:
                                                        v2 = append(v2, fi35)
                                                    case 42:
                                                        v63 -= -1
                                                        v70 := v4[v64]
                                                        v39 = v70
                                                        _ = v70
                                                    default:
                                                        var v71 bool = true
                                                        _ = v71
                                                    }
                                                    if (127 <= (v67 * 100)) {
                                                        var v72 int32 = (-(v63))
                                                        v1 += (v1 - v1)
                                                        v64 -= (v64 + v64)
                                                        var v73 int8 = (v67 + 42)
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    _ = v69
                                                }
                                            }
                                            fmt.Println(v67)
                                            if ((-(v63)) >= (v63 * v63)) {
                                                var v74 bool = (((0.1 - 0.0) - float64((float32((((3.14159 - 3.14159) / float64(1.5)) / 3.14159)) + 0.0))) <= -2.25)
                                                _ = v74
                                            }
                                            delete(v3, v64)
                                            var v75 float64 = float64(((1.5 - (0.0 - (1e10 - 1.5))) * 0.1))
                                            v39 = (^(v39))
                                            _ = v75
                                        } else {
                                            var v76 int = fi61
                                            switch "" {
                                            case "foo bar":
                                                v77 := GenS7{F0: (fi35 <= 0), F1: v63}
                                                _ = v77
                                            default:
                                                var v78 int32 = -2147483648
                                                _ = v78
                                            }
                                            _ = v76
                                        }
                                        v79 := GenS6{F0: "a", F1: (fi35 ^ (fi61 ^ 1)), F2: int8(v64)}
                                        v80 := GenS8{F0: "a", F1: v1}
                                        _ = v68
                                        _ = v79
                                        _ = v80
                                    }
                                    _ = v63
                                    _ = v64
                                    _ = v66
                                    _ = v67
                                }
                                _ = fi61
                            }
                        case 128:
                            v81 := GenS6{F0: ("\n\t" + "a"), F1: 0, F2: -1}
                            _ = v81
                        case 0:
                            v1 = v1
                        default:
                            v4[2147483647] = v39
                        }
                    }
                } else {
                    switch ("hello" + (("\n\t" + (("a" + "foo bar") + ("hello" + "hello"))) + "hello")) {
                    case "hello":
                        v39 = (v39 >> ((v1) & 63))
                    case "a":
                        switch "\n\t" {
                        case "hello":
                            var v82 int32 = -322
                            _ = v82
                        case "a":
                            var v83 float64 = (0.1 - (1e10 - 1.5))
                            _ = v83
                        case "\n\t":
                            v84 := GenS3{F0: 9223372036854775807, F1: 18446744073709551615}
                            _ = v84
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    default:
                        delete(v3, 65535)
                    }
                    switch (fi35 >> ((v1) & 63)) {
                    case 234:
                        if (!true) {
                            if false {
                                fmt.Println(fi35, fi35)
                                for fi85 := 0; fi85 < 19; fi85++ {
                                    var v86 uint16 = 152
                                    v87 := GenS9{F0: 964, F1: 1}
                                    _ = fi85
                                    _ = v86
                                    _ = v87
                                }
                                var v88 uint16 = 1
                                _ = v88
                            }
                            for fi89 := 0; fi89 < 6; fi89++ {
                                switch "foo bar" {
                                case "\n\t":
                                    v4[0] = (fi89 + (fi89 | -912))
                                }
                                v1 += v1
                                _ = fi89
                            }
                            v90 := GenS2{F0: 394, F1: v1, F2: 255}
                            _ = v90
                        }
                    case 32767:
                        for fi91 := 0; fi91 < 18; fi91++ {
                            if (0 <= 2) {
                                if len(v5) > 0 {
                                    v5[0] = v39
                                }
                                if ((1.5 > 0.0) && true) {
                                    var v92 uint64 = 195
                                    v93 := GenS10{F0: (v92 / ((v92) | 1)), F1: 1}
                                    _ = v92
                                    _ = v93
                                } else {
                                    v94 := GenS10{F0: 4294967295, F1: 127}
                                    var v95 uint = 888
                                    _ = v94
                                    _ = v95
                                }
                            }
                            if false {
                                if true {
                                    v4[909] = int(v1)
                                    v96 := GenS11{F0: 146}
                                    v3[0] = false
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v96
                                }
                                v39 = -626
                                v4[146] = fi91
                            } else {
                                var v97 uint64 = 42
                                if (true || false) {
                                    if ((v39 + (-(v39))) != fi35) {
                                        var v98 uint64 = (v97 + v97)
                                        var v99 int = 4294967295
                                        v97 = uint64(v1)
                                        _ = v98
                                        _ = v99
                                    }
                                    if (!(-74 >= 6)) {
                                        fmt.Println(v39, v1, fi91)
                                    } else {
                                        if (true && (((42 < -1) || (true || false)) || ((float32(3.14159) + 0.0) <= 3.14159))) {
                                            switch "a" {
                                            case "\n\t":
                                                v100 := GenS12{F0: false, F1: 1}
                                                _ = v100
                                            case "foo bar":
                                                switch (fi35 + (v39 & 128)) {
                                                case 4294967295:
                                                    v1 = v1
                                                case 0:
                                                    fmt.Println(v39, v97, v39)
                                                    var v101 uint64 = 2
                                                    _ = v101
                                                }
                                            case "a":
                                                v102 := GenS8{F0: "", F1: v1}
                                                _ = v102
                                            }
                                        }
                                        if true {
                                            var v103 float32 = ((1e10 + float32((1.5 - (1e10 / 1.0)))) * ((1.5 - 0.0) - float32(float64((3.14159 / 1e10)))))
                                            v1 += v1
                                            _ = v103
                                        }
                                    }
                                    fmt.Println(v39)
                                } else {
                                    if true {
                                        v104 := GenS13{F0: "foo bar", F1: -98, F2: (1 == 2147483647)}
                                        if len(v2) > 0 {
                                            v2[0] = (fi91 & fi35)
                                        }
                                        _ = v104
                                    }
                                }
                                var v105 float64 = (3.14159 * 0.0)
                                _ = v105
                                _ = v97
                            }
                            var v106 int64 = int64(v39)
                            fmt.Println(v39)
                            _ = fi91
                            _ = v106
                        }
                    case 987:
                        var v107 uint32 = 0
                        _ = v107
                    default:
                        if true {
                            fmt.Println(v1)
                            for fi108 := 0; fi108 < 3; fi108++ {
                                var v109 int32 = 0
                                v39 = v39
                                _ = fi108
                                _ = v109
                            }
                        } else {
                            v110 := GenS2{F0: 4294967295, F1: 255, F2: 127}
                            _ = v110
                        }
                    }
                }
            case "foo bar":
                v1 += v1
            default:
                v4[1] = v39
            }
            switch 32767 {
            case -2147483648:
                switch -474 {
                case 42:
                    switch "" {
                    case "":
                        v39 = fi35
                    case "hello":
                        v111 := GenS3{F0: 0, F1: 536}
                        _ = v111
                    case "\n\t":
                        for fi112 := 0; fi112 < 10; fi112++ {
                            var v113 int8 = -1
                            _ = fi112
                            _ = v113
                        }
                    }
                case 607:
                    fmt.Println(fi35)
                case 4294967295:
                    for fi114 := 0; fi114 < 10; fi114++ {
                        var v115 int16 = int16(v39)
                        v115 = (-(v115))
                        _ = fi114
                        _ = v115
                    }
                }
            case 128:
                fmt.Println(fi35, v1, v39)
            case -176:
                v39 = 4294967295
            }
            _ = v39
        }
        _ = fi35
    }
    var v116 int32 = 2147483647
    v2 = append(v2, int(v116))
    delete(v4, 1)
    v117 := GenS1{F0: 0}
    v3[2147483647] = (-9223372036854775808 == 1)
    var v118 uint64 = 0
    v118 -= v118
    if (239 < 255) {
        fmt.Println(v118, v1)
    } else {
        var v119 float32 = float32(0.0)
        if true {
            v118 -= v118
            v120 := GenS14{F0: 105, F1: 0, F2: "\n\t"}
            switch "hello" {
            case "hello":
                v121 := GenS6{F0: ("a" + ("hello" + "\n\t")), F1: 42, F2: 19}
                _ = v121
            case "\n\t":
                if true {
                    var v122 uint16 = 65535
                    _ = v122
                } else {
                    fmt.Println(v119, v118)
                    var v123 uint64 = 1
                    _ = v123
                }
            }
            _ = v120
        } else {
            v119 = v119
        }
        v124 := GenS15{F0: v1, F1: 192}
        _ = v119
        _ = v124
    }
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v118:", v118)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v116
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
